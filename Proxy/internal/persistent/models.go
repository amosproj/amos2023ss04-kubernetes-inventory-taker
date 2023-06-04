// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Cluster struct {
	ClusterEventID int32            `db:"cluster_event_id" json:"clusterEventID"`
	ClusterID      int32            `db:"cluster_id" json:"clusterID"`
	Timestamp      pgtype.Timestamp `db:"timestamp" json:"timestamp"`
	Name           string           `db:"name" json:"name"`
}

type Container struct {
	ContainerEventID int32            `db:"container_event_id" json:"containerEventID"`
	ContainerID      pgtype.Int4      `db:"container_id" json:"containerID"`
	Timestamp        pgtype.Timestamp `db:"timestamp" json:"timestamp"`
	PodID            pgtype.UUID      `db:"pod_id" json:"podID"`
	Name             pgtype.Text      `db:"name" json:"name"`
	Image            pgtype.Text      `db:"image" json:"image"`
	Status           pgtype.Text      `db:"status" json:"status"`
	Ports            pgtype.Text      `db:"ports" json:"ports"`
}

type Node struct {
	NodeEventID                      int32            `db:"node_event_id" json:"nodeEventID"`
	NodeID                           pgtype.UUID      `db:"node_id" json:"nodeID"`
	Timestamp                        pgtype.Timestamp `db:"timestamp" json:"timestamp"`
	CreationTime                     pgtype.Timestamp `db:"creation_time" json:"creationTime"`
	Name                             pgtype.Text      `db:"name" json:"name"`
	IpAddressInternal                []string         `db:"ip_address_internal" json:"ipAddressInternal"`
	IpAddressExternal                []string         `db:"ip_address_external" json:"ipAddressExternal"`
	Hostname                         pgtype.Text      `db:"hostname" json:"hostname"`
	StatusCapacityCpu                pgtype.Text      `db:"status_capacity_cpu" json:"statusCapacityCpu"`
	StatusCapacityMemory             pgtype.Text      `db:"status_capacity_memory" json:"statusCapacityMemory"`
	StatusCapacityPods               pgtype.Text      `db:"status_capacity_pods" json:"statusCapacityPods"`
	StatusAllocatableCpu             pgtype.Text      `db:"status_allocatable_cpu" json:"statusAllocatableCpu"`
	StatusAllocatableMemory          pgtype.Text      `db:"status_allocatable_memory" json:"statusAllocatableMemory"`
	StatusAllocatablePods            pgtype.Text      `db:"status_allocatable_pods" json:"statusAllocatablePods"`
	KubeletVersion                   pgtype.Text      `db:"kubelet_version" json:"kubeletVersion"`
	NodeConditionsReady              pgtype.Text      `db:"node_conditions_ready" json:"nodeConditionsReady"`
	NodeConditionsDiskPressure       pgtype.Text      `db:"node_conditions_disk_pressure" json:"nodeConditionsDiskPressure"`
	NodeConditionsMemoryPressure     pgtype.Text      `db:"node_conditions_memory_pressure" json:"nodeConditionsMemoryPressure"`
	NodeConditionsPidPressure        pgtype.Text      `db:"node_conditions_pid_Pressure" json:"nodeConditionsPidPressure"`
	NodeConditionsNetworkUnavailable pgtype.Text      `db:"node_conditions_network_unavailable" json:"nodeConditionsNetworkUnavailable"`
}

type Pod struct {
	PodResourceVersion int32            `db:"pod_resource_version" json:"podResourceVersion"`
	PodID              pgtype.UUID      `db:"pod_id" json:"podID"`
	Timestamp          pgtype.Timestamp `db:"timestamp" json:"timestamp"`
	ClusterID          pgtype.Int4      `db:"cluster_id" json:"clusterID"`
	NodeID             pgtype.Text      `db:"node_id" json:"nodeID"`
	Name               pgtype.Text      `db:"name" json:"name"`
	Namespace          pgtype.Text      `db:"namespace" json:"namespace"`
	Status             pgtype.Text      `db:"status" json:"status"`
	Stamp2             pgtype.Timestamp `db:"stamp2" json:"stamp2"`
}

type Service struct {
	Name              string           `db:"name" json:"name"`
	Namespace         string           `db:"namespace" json:"namespace"`
	Timestamp         pgtype.Timestamp `db:"timestamp" json:"timestamp"`
	Labels            []string         `db:"labels" json:"labels"`
	CreationTimestamp pgtype.Timestamp `db:"creation_timestamp" json:"creationTimestamp"`
	Ports             []string         `db:"ports" json:"ports"`
	ExternalIps       []string         `db:"external_ips" json:"externalIps"`
	ClusterIp         string           `db:"cluster_ip" json:"clusterIp"`
}
