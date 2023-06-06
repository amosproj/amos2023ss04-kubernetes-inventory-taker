package cluster

import (
	"context"
	"encoding/json"
	"log"

	model "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog"
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

func ProcessNode(event Event, bunDB *bun.DB) {
	//nolint:forcetypeassert
	nodeNew := event.Object.(*corev1.Node)

	//nolint:forcetypeassert
	if event.Type == Update && event.OldObj.(*corev1.Node).ResourceVersion == nodeNew.ResourceVersion {
		return
	}

	conditions := getNodeConditions(nodeNew)
	capacity, allocatable := getStatus(nodeNew)
	internalIPs, externalIPs, hostname := getNodeAddresses(nodeNew)

	jsonData, err := json.Marshal(nodeNew)
	if err != nil {
		klog.Error("Error converting Node to JSON:", err)
		return
	}

	// Create a new Node
	nodeDB := model.Node{
		BaseModel: schema.BaseModel{},

		NodeID:                  string(nodeNew.UID),
		Timestamp:               event.timestamp,
		CreationTime:            nodeNew.CreationTimestamp.Time,
		Name:                    nodeNew.Name,
		IPAddressInternal:       internalIPs,
		IPAddressExternal:       externalIPs,
		Hostname:                hostname,
		StatusCapacityCPU:       capacity.CPU.String(),
		StatusCapacityMemory:    capacity.Memory.String(),
		StatusCapacityPods:      capacity.Pods.String(),
		StatusAllocatableCPU:    allocatable.CPU.String(),
		StatusAllocatableMemory: allocatable.Memory.String(),
		StatusAllocatablePods:   allocatable.Pods.String(),
		KubeletVersion:          nodeNew.Status.NodeInfo.KubeletVersion,
		Ready:                   conditions.Ready,
		DiskPressure:            conditions.DiskPressure,
		MemoryPressure:          conditions.MemoryPressure,
		PIDPressure:             conditions.PIDPressure,
		NetworkUnavailable:      conditions.NetworkUnavailable,
		Data:                    string(jsonData),
	}

	// Insert the new Node into the database
	_, err = bunDB.NewInsert().Model(&nodeDB).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

// getNodeAddresses separates the internal and external IP addresses and the hostname from a node's status.
func getNodeAddresses(node *corev1.Node) ([]string, []string, string) {
	var internalIPs, externalIPs []string

	var hostname string

	for _, address := range node.Status.Addresses {
		switch address.Type {
		case corev1.NodeInternalIP:
			internalIPs = append(internalIPs, address.Address)
		case corev1.NodeExternalIP:
			externalIPs = append(externalIPs, address.Address)
		case corev1.NodeHostName:
			hostname = address.Address
		case corev1.NodeExternalDNS:
		case corev1.NodeInternalDNS:
		}
	}

	return internalIPs, externalIPs, hostname
}

func getNodeConditions(node *corev1.Node) NodeConditions {
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

// Does return status of cpu, memory and pods; allocatable are resources of a node that are available for scheduling.
func getStatus(node *corev1.Node) (Capacity, Allocatable) {
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
