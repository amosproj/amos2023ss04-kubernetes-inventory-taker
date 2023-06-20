package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Cluster struct {
	bun.BaseModel `bun:"table:Clusters"`
	ID            int       `bun:"id,autoincrement"`
	ClusterID     int       `bun:"cluster_id"`
	Timestamp     time.Time `bun:"timestamp,notnull"`
	Name          string    `bun:"name"`
}

type Service struct {
	bun.BaseModel     `bun:"Services"`
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
	bun.BaseModel      `bun:"Pods"`
	ID                 int       `bun:"id,autoincrement"`
	PodResourceVersion string    `bun:"pod_resource_version,type:text,notnull"`
	PodID              string    `bun:"pod_id,type:uuid,notnull"`
	Timestamp          time.Time `bun:"timestamp,type:timestamp,notnull"`
	NodeName           string    `bun:"node_name,type:text"`
	Name               string    `bun:"name,type:text"`
	Namespace          string    `bun:"namespace,type:text"`
	StatusPhase        string    `bun:"status_phase,type:text"`
	Data               string    `bun:"data,type:json"`
}

type Node struct {
	bun.BaseModel           `bun:"table:Nodes"`
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
	bun.BaseModel `bun:"table:containers"`
	ID            int             `bun:"id,autoincrement,pk"`
	Timestamp     time.Time       `bun:"timestamp,type:timestamp,notnull"`
	ContainerID   string          `bun:"container_id,type:text"`
	PodID         string          `bun:"pod_id,type:uuid"`
	Name          string          `bun:"name,type:text"`
	Image         string          `bun:"image,type:text"`
	Status        string          `bun:"status,type:text"`
	Ports         string          `bun:"ports,type:text"`
	ImageID       string          `bun:"image_id,type:text"`
	Ready         bool            `bun:"ready"`
	RestartCount  int             `bun:"restart_count"`
	Started       bool            `bun:"started"`
	StateID       int             `bun:"state_id"`
	LastStateID   int             `bun:"last_state_id,nullzero"`
	State         *ContainerState `bun:"rel:belongs-to,join:state=id"`
	LastState     *ContainerState `bun:"rel:belongs-to,join:last_state=id"`
}

type ContainerState struct {
	bun.BaseModel `bun:"table:container_states"`
	ID            int       `bun:"id,autoincrement,pk"`
	Kind          string    `bun:"kind,type:text"`
	StartedAt     time.Time `bun:"started_at,type:time"`
	ContainerID   string    `bun:"container_id,type:text"`
	ExitCode      int       `bun:"exit_code,type:int"`
	FinishedAt    time.Time `bun:"finished_at,type:time"`
	Message       string    `bun:"message,type:text"`
	Reason        string    `bun:"reason,type:text"`
	Signal        int       `bun:"signal,type:int"`
}

type VolumeDevice struct {
	bun.BaseModel `bun:"table:volume_devices"`
	ID            int    `bun:"id,autoincrement,pk"`
	ContainerID   int    `bun:"container_id,type:int"`
	DevicePath    string `bun:"device_path,type:text"`
	Name          string `bun:"name,type:text"`
}

type VolumeMount struct {
	bun.BaseModel    `bun:"table:volume_mounts"`
	ID               int    `bun:"id,autoincrement,pk"`
	ContainerID      int    `bun:"container_id,type:int"`
	MountPath        string `bun:"mount_path,type:text"`
	MountPropagation string `bun:"mount_propagation,type:text"`
	Name             string `bun:"name,type:text"`
	ReadOnly         bool   `bun:"read_only,type:bool"`
	SubPath          string
	SubPathExpr      string
}

type ContainerPort struct {
	bun.BaseModel `bun:"table:container_ports"`
	ID            int    `bun:"id,autoincrement,pk"`
	ContainerID   int    `bun:"container_id,type:int"`
	ContainerPort int    `bun:"container_port,type:int"`
	HostIP        string `bun:"host_ip,type:text"`
	HostPort      int    `bun:"host_port,type:int"`
	Name          string `bun:"name,type:text"`
	Protocol      string `bun:"protocol,type:text"`
}
