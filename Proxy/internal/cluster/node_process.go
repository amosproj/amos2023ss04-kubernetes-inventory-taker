package cluster

import (
	"context"
	"math/rand"

	database "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/persistent"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
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

func ProcessNode(event Event, dbQueries *database.Queries) {
	nodeNew, assertOk := event.Object.(*corev1.Node)
	if !assertOk {
		klog.Error(nil)
	}

	nodeOld, assertOk := event.OldObj.(*corev1.Node)
	if !assertOk {
		klog.Error(nil)
	}

	if event.Type == Update && nodeOld.ResourceVersion == nodeNew.ResourceVersion {
		return
	}

	conditions := GetNodeConditions(nodeNew)
	capacity, allocatable := GetStatus(nodeNew)
	internalIPs, externalIPs, hostname := GetNodeAddresses(nodeNew)

	// Create a new Node
	var nodeDB database.UpdateNodeParams

	nodeDB.NodeEventID = int32(rand.Int31())
	nodeDB.NodeID.Scan(nodeNew.UID)
	nodeDB.Timestamp.Scan(event.timestamp)
	nodeDB.CreationTime.Scan(nodeNew.CreationTimestamp.Time)
	nodeDB.Name.Scan(nodeNew.Name)
	nodeDB.IpAddressInternal = internalIPs
	nodeDB.IpAddressExternal = externalIPs
	nodeDB.Hostname.Scan(hostname)
	nodeDB.StatusCapacityCpu.Scan(capacity.CPU.String())
	nodeDB.StatusCapacityMemory.Scan(capacity.Memory.String())
	nodeDB.StatusCapacityPods.Scan(capacity.Pods.String())
	nodeDB.StatusAllocatableCpu.Scan(allocatable.CPU.String())
	nodeDB.StatusAllocatableMemory.Scan(allocatable.Memory.String())
	nodeDB.StatusAllocatablePods.Scan(allocatable.Pods.String())
	nodeDB.KubeletVersion.Scan(nodeNew.Status.NodeInfo.KubeletVersion)
	nodeDB.NodeConditionsReady.Scan(conditions.Ready)
	nodeDB.NodeConditionsDiskPressure.Scan(conditions.DiskPressure)
	nodeDB.NodeConditionsMemoryPressure.Scan(conditions.MemoryPressure)
	nodeDB.NodeConditionsPidPressure.Scan(conditions.PIDPressure)
	nodeDB.NodeConditionsNetworkUnavailable.Scan(conditions.NetworkUnavailable)

	// Insert the new Node into the database
	err := dbQueries.UpdateNode(context.Background(), nodeDB)
	if err != nil {
		klog.Fatal(err)
	}
}

// GetNodeAddresses separates the internal and external IP addresses and the hostname from a node's status.
func GetNodeAddresses(node *corev1.Node) ([]string, []string, string) {
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
		case corev1.NodeInternalDNS:
		case corev1.NodeExternalDNS:
			klog.Info("Nodes external DNS", "IP", address.Address)
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

// Does return status of cpu, memory and pods; allocatable are resources of a node that are available for scheduling.
func GetStatus(node *corev1.Node) (Capacity, Allocatable) {
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
