// Package model contains all the bun models for the database.
package model

import (
	"time"
)

type Service struct {
	ID                int       `bun:"id,autoincrement"`
	Name              string    `bun:"name,type:text,notnull,pk"`
	Namespace         string    `bun:"namespace,type:text,notnull,pk"`
	Timestamp         time.Time `bun:"timestamp,type:timestamp,notnull"`
	Labels            []string  `bun:"labels,type:text[],array,notnull"`
	CreationTimestamp time.Time `bun:"creation_timestamp,type:timestamp,notnull"`
	Ports             []string  `bun:"ports,type:text[],array,notnull"`
	ExternalIPs       []string  `bun:"external_ips,type:text[],array"`
	ClusterIP         string    `bun:"cluster_ip,type:text,notnull"`
	Data              string    `bun:"data,type:json"`
}

type Pod struct {
	ID                 int       `bun:"id,autoincrement,pk"`
	PodResourceVersion string    `bun:"pod_resource_version,type:text,notnull"`
	PodID              string    `bun:"pod_id,type:uuid,notnull"`
	Timestamp          time.Time `bun:"timestamp,type:timestamp,notnull"`
	NodeName           string    `bun:"node_name,type:text"`
	Name               string    `bun:"name,type:text"`
	Namespace          string    `bun:"namespace,type:text"`
	StatusPhase        string    `bun:"status_phase,type:text"`
	Data               string    `bun:"data,type:json"`
	HostIP             string    `bun:"host_ip"`
	PodIP              string    `bun:"pod_ip"`
	PodIPs             []string  `bun:"pod_ips,array"`
	StartTime          time.Time `bun:"start_time"`
	QOSClass           string    `bun:"qos_class"`
}

type PodStatusCondition struct {
	ID                 int       `bun:"id,autoincrement,notnull,pk"`
	PodID              int       `bun:"pod_id"`
	Type               string    `bun:"type"`
	Status             string    `bun:"status"`
	LastProbeTime      time.Time `bun:"last_probe_time,type:timestamp,nullzero"`
	LastTransitionTime time.Time `bun:"last_transition_time,type:timestamp"`
	Reason             string    `bun:"reason"`
	Message            string    `bun:"message"`
}

type PodVolume struct {
	ID        int    `bun:"id,autoincrement,notnull,pk"`
	PodID     int    `bun:"pod_id"`
	Type      string `bun:"type"`
	Name      string `bun:"name"`
	ClaimName string `bun:"persistent_claim_name"`
	ReadOnly  bool   `bun:"read_only"`
}

type Node struct {
	ID                      int       `bun:"id,autoincrement"`
	NodeID                  string    `bun:"node_id,type:uuid"`
	Timestamp               time.Time `bun:"timestamp,type:timestamp,notnull"`
	CreationTime            time.Time `bun:"creation_time,type:timestamp"`
	Name                    string    `bun:"name,type:text"`
	IPAddressInternal       []string  `bun:"ip_address_internal,type:text[],array"`
	IPAddressExternal       []string  `bun:"ip_address_external,type:text[],array"`
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
	Data                    string    `bun:"data,type:json"`
}

type Container struct {
	ID           int       `bun:"id,autoincrement,pk"`
	Timestamp    time.Time `bun:"timestamp,type:timestamp,notnull"`
	ContainerID  string    `bun:"container_id,type:text"`
	PodID        string    `bun:"pod_id,type:uuid"`
	Name         string    `bun:"name,type:text"`
	Image        string    `bun:"image,type:text"`
	Status       string    `bun:"status,type:text"`
	Ports        string    `bun:"ports,type:text"`
	ImageID      string    `bun:"image_id,type:text"`
	Ready        bool      `bun:"ready"`
	RestartCount int       `bun:"restart_count"`
	Started      bool      `bun:"started"`
	StateID      int       `bun:"state_id"`
	// since LastState can be unset, it should automatically me NULL in the database instead of 0
	LastStateID int `bun:"last_state_id,nullzero"`
	// theses references are not used for inserting, as bun does not support that
	State     *ContainerState `bun:"rel:belongs-to,join:state=id"`
	LastState *ContainerState `bun:"rel:belongs-to,join:last_state=id"`
}

type ContainerState struct {
	ID          int       `bun:"id,autoincrement,pk"`
	Kind        string    `bun:"kind,type:text"`
	StartedAt   time.Time `bun:"started_at,type:time"`
	ContainerID string    `bun:"container_id,type:text"`
	ExitCode    int       `bun:"exit_code,type:int"`
	FinishedAt  time.Time `bun:"finished_at,type:time"`
	Message     string    `bun:"message,type:text"`
	Reason      string    `bun:"reason,type:text"`
	Signal      int       `bun:"signal,type:int"`
}

type VolumeDevice struct {
	ID          int    `bun:"id,autoincrement,pk"`
	ContainerID int    `bun:"container_id,type:int"`
	DevicePath  string `bun:"device_path,type:text"`
	Name        string `bun:"name,type:text"`
}

type VolumeMount struct {
	ID               int    `bun:"id,autoincrement,pk"`
	ContainerID      int    `bun:"container_id,type:int"`
	MountPath        string `bun:"mount_path,type:text"`
	MountPropagation string `bun:"mount_propagation,type:text"`
	Name             string `bun:"name,type:text"`
	ReadOnly         bool   `bun:"read_only,type:bool"`
	SubPath          string `bun:"sub_path,type:text"`
	SubPathExpr      string `bun:"sub_path_expr,type:text"`
}

type ContainerPort struct {
	ID            int    `bun:"id,autoincrement,pk"`
	ContainerID   int    `bun:"container_id,type:int"`
	ContainerPort int    `bun:"container_port,type:int"`
	HostIP        string `bun:"host_ip,type:text"`
	HostPort      int    `bun:"host_port,type:int"`
	Name          string `bun:"name,type:text"`
	Protocol      string `bun:"protocol,type:text"`
}

type PersistentVolume struct {
	ID                int       `bun:"id,autoincrement,pk"`
	Timestamp         time.Time `bun:"timestamp,type:timestamp,notnull"`
	Labels            []string  `bun:"labels,type:text[],array"`
	Name              string    `bun:"name,type:text,notnull"`
	Namespace         string    `bun:"namespace,type:text,notnull"`
	UID               string    `bun:"uid,type:text"`
	CreationTimestamp time.Time `bun:"creation_timestamp,type:timestamp"`
	DeletionTimestamp time.Time `bun:"deletion_timestamp,type:timestamp"`
	AccessModes       []string  `bun:"access_modes,type:text[],array"`
	Capacity          string    `bun:"capacity,type:text"`
	MountOptions      []string  `bun:"mount_options,type:text[],array"`
	StorageClassName  string    `bun:"storage_class_name,type:text"`
	VolumeMode        string    `bun:"volume_mode,type:text"`
	Message           string    `bun:"message,type:text"`
	Phase             string    `bun:"phase,type:text"`
	Reason            string    `bun:"reason,type:text"`
}

type PersistentVolumeClaim struct {
	ID                int       `bun:"id,autoincrement,pk"`
	Timestamp         time.Time `bun:"timestamp,type:timestamp,notnull"`
	Labels            []string  `bun:"labels,type:text[],array"`
	Name              string    `bun:"name,type:text,notnull"`
	Namespace         string    `bun:"namespace,type:text,notnull"`
	UID               string    `bun:"uid,type:text"`
	CreationTimestamp time.Time `bun:"creation_timestamp,type:timestamp"`
	DeletionTimestamp time.Time `bun:"deletion_timestamp,type:timestamp"`
	AccessModes       []string  `bun:"access_modes,type:text[],array"`
	StorageClassName  string    `bun:"storage_class_name,type:text"`
	VolumeMode        string    `bun:"volume_mode,type:text"`
	VolumeName        string    `bun:"volume_name,type:text"`
	Capacity          string    `bun:"capacity,type:text"`
	Phase             string    `bun:"phase,type:text"`
	ResizeStatus      string    `bun:"resize_status,type:text"`
}

type PersistentVolumeClaimCondition struct {
	ID                      int       `bun:"id,autoincrement,pk"`
	PersistentVolumeClaimID int       `bun:"persistent_volume_claim_id,type:int"`
	LastProbeTime           time.Time `bun:"last_probe_time,type:timestamp"`
	LastTransitionTime      time.Time `bun:"last_transition_time,type:timestamp"`
	Message                 string    `bun:"message,type:text"`
	Reason                  string    `bun:"reason,type:text"`
	Status                  string    `bun:"status,type:text"`
	Type                    string    `bun:"type,type:text"`
}
