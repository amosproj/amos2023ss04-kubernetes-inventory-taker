package cluster

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type NodeConditions struct {
	Ready              string
	DiskPressure       string
	MemoryPressure     string
	PIDPressure        string
	NetworkUnavailable string
}

type Capacity struct {
	CPU    resource.Quantity
	Memory resource.Quantity
	Pods   resource.Quantity
}

type Allocatable struct {
	CPU    resource.Quantity
	Memory resource.Quantity
	Pods   resource.Quantity
}

type Node struct {
	bun.BaseModel           `bun:"table:Node"`
	NodeEventID             int       `bun:"node_event_id,type:integer,notnull"`
	NodeID                  string    `bun:"node_id,type:uuid"`
	Timestamp               time.Time `bun:"timestamp,type:timestamp,notnull"`
	CreationTime            time.Time `bun:"creation_time,type:timestamp"`
	Name                    string    `bun:"name,type:text"`
	IPAddressInternal       []string  `bun:"ip_address_internal,type:varchar[],array"`
	IPAddressExternal       []string  `bun:"ip_address_external,type:varchar[],array"`
	Hostname                string    `bun:"hostname,type:text"`
	StatusCapacityCPU       string    `bun:"status_capacity_cpu,type:text"`
	StatusCapacityMemory    string    `bun:"status_capacity_memory,type:text"`
	StatusCapacityPods      string    `bun:"status_capacity_pods,type:text"`
	StatusAllocatableCPU    string    `bun:"status_allocatable_cpu,type:text"`
	StatusAllocatableMemory string    `bun:"status_allocatable_memory,type:text"`
	StatusAllocatablePods   string    `bun:"status_allocatable_pods,type:text"`
	KubeletVersion          string    `bun:"kubelet_version,type:text"`
	Ready                   string    `bun:"node_conditions_ready,type:text"`
	DiskPressure            string    `bun:"node_conditions_disk_pressure,type:text"`
	MemoryPressure          string    `bun:"node_conditions_memory_pressure,type:text"`
	PIDPressure             string    `bun:"node_conditions_pid_Pressure,type:text"`
	NetworkUnavailable      string    `bun:"node_conditions_network_unavailable,type:text"`
}

func ProcessNode(event Event, db *bun.DB) {
	node := event.Object.(*corev1.Node)

	//ignore update event with unchanged resource version
	if event.Type == Update && event.OldObj.(*corev1.Node).ResourceVersion == node.ResourceVersion {
		return
	}

	//get some "more complicated to obtain" fields
	conditions := GetNodeConditions(node)
	capacity, allocatable := GetStatus(node)
	internalIPs, externalIPs, hostname := GetNodeAddresses(node)

	// Create a new Node
	nodeDB := Node{
		BaseModel:               schema.BaseModel{},
		NodeEventID:             int(rand.Int31()),
		NodeID:                  string(node.UID),
		Timestamp:               event.timestamp,
		CreationTime:            node.CreationTimestamp.Time,
		Name:                    node.Name,
		IPAddressInternal:       internalIPs,
		IPAddressExternal:       externalIPs,
		Hostname:                hostname,
		StatusCapacityCPU:       capacity.CPU.String(),
		StatusCapacityMemory:    capacity.Memory.String(),
		StatusCapacityPods:      capacity.Pods.String(),
		StatusAllocatableCPU:    allocatable.CPU.String(),
		StatusAllocatableMemory: allocatable.Memory.String(),
		StatusAllocatablePods:   allocatable.Pods.String(),
		KubeletVersion:          node.Status.NodeInfo.KubeletVersion,
		Ready:                   conditions.Ready,
		DiskPressure:            conditions.DiskPressure,
		MemoryPressure:          conditions.MemoryPressure,
		PIDPressure:             conditions.PIDPressure,
		NetworkUnavailable:      conditions.NetworkUnavailable,
	}

	// Insert the new Node into the database
	_, err := db.NewInsert().Model(&nodeDB).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

// GetNodeAddresses separates the internal and external IP addresses and the hostname from a node's status.
func GetNodeAddresses(node *corev1.Node) (internalIPs []string, externalIPs []string, hostname string) {
	for _, address := range node.Status.Addresses {
		switch address.Type {
		case corev1.NodeInternalIP:
			internalIPs = append(internalIPs, address.Address)
		case corev1.NodeExternalIP:
			externalIPs = append(externalIPs, address.Address)
		case corev1.NodeHostName:
			hostname = address.Address
		}
	}
	return internalIPs, externalIPs, hostname
}

func GetNodeConditions(node *corev1.Node) NodeConditions {
	// Initialize NodeConditions with default status
	conditions := NodeConditions{
		Ready:              "Unknown",
		DiskPressure:       "Unknown",
		MemoryPressure:     "Unknown",
		PIDPressure:        "Unknown",
		NetworkUnavailable: "Unknown",
	}

	// Iterate over the node's conditions and update the appropriate status
	for _, condition := range node.Status.Conditions {
		switch condition.Type {
		case corev1.NodeReady:
			conditions.Ready = string(condition.Status)
		case corev1.NodeDiskPressure:
			conditions.DiskPressure = string(condition.Status)
		case corev1.NodeMemoryPressure:
			conditions.MemoryPressure = string(condition.Status)
		case corev1.NodePIDPressure:
			conditions.PIDPressure = string(condition.Status)
		case corev1.NodeNetworkUnavailable:
			conditions.NetworkUnavailable = string(condition.Status)
		}
	}

	return conditions
}

// Does return status of cpu, memory and pods; allocatable are resources of a node that are available for scheduling
func GetStatus(node *corev1.Node) (capacity Capacity, allocatable Allocatable) {
	return Capacity{
			CPU:    node.Status.Capacity[corev1.ResourceCPU],
			Memory: node.Status.Capacity[corev1.ResourceMemory],
			Pods:   node.Status.Capacity[corev1.ResourcePods],
		},
		Allocatable{
			CPU:    node.Status.Allocatable[corev1.ResourceCPU],
			Memory: node.Status.Allocatable[corev1.ResourceMemory],
			Pods:   node.Status.Allocatable[corev1.ResourcePods],
		}
}
