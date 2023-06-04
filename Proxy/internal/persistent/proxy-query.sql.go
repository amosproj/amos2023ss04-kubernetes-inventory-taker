// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: proxy-query.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const updateCluster = `-- name: UpdateCluster :exec
INSERT INTO "Cluster" (
    "cluster_id",
    "timestamp",
    "name"
  )
VALUES ($1, $2, $3)
`

type UpdateClusterParams struct {
	ClusterID int32            `db:"cluster_id" json:"clusterID"`
	Timestamp pgtype.Timestamp `db:"timestamp" json:"timestamp"`
	Name      string           `db:"name" json:"name"`
}

func (q *Queries) UpdateCluster(ctx context.Context, arg UpdateClusterParams) error {
	_, err := q.db.Exec(ctx, updateCluster, arg.ClusterID, arg.Timestamp, arg.Name)
	return err
}

const updateNode = `-- name: UpdateNode :exec
INSERT INTO "Node" (
    "node_id",
    "timestamp",
    "creation_time",
    "name",
    "ip_address_internal",
    "ip_address_external",
    "hostname",
    "status_capacity_cpu",
    "status_capacity_memory",
    "status_capacity_pods",
    "status_allocatable_cpu",
    "status_allocatable_memory",
    "status_allocatable_pods",
    "kubelet_version",
    "node_conditions_ready",
    "node_conditions_disk_pressure",
    "node_conditions_memory_pressure",
    "node_conditions_pid_Pressure",
    "node_conditions_network_unavailable"
  )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15,
    $16,
    $17,
    $18,
    $19
  )
`

type UpdateNodeParams struct {
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

// -- name: UpdatePod :one
// INSERT INTO "Pod" (
//
//	"pod_resource_version", "pod_id", "timestamp", "cluster_id", "node_id", "name", "namespace","status", "stamp2"
//
// ) VALUES (
//
//	$1, $2, $3, $4, $5, $6, $7, $8, $9
//
// )
// RETURNING "pod_resource_version", "pod_id", "timestamp", "cluster_id", "node_id", "name", "namespace","status", "stamp2";
//
// -- name: UpdateContainer :one
// INSERT INTO "Container" (
//
//	"container_event_id", "container_id", "timestamp", "pod_id", "name", "image", "status", "ports"
//
// ) VALUES (
//
//	$1, $2, $3, $4, $5, $6, $7, $8
//
// )
// RETURNING "container_event_id", "container_id", "timestamp", "pod_id", "name", "image", "status", "ports";
func (q *Queries) UpdateNode(ctx context.Context, arg UpdateNodeParams) error {
	_, err := q.db.Exec(ctx, updateNode,
		arg.NodeID,
		arg.Timestamp,
		arg.CreationTime,
		arg.Name,
		arg.IpAddressInternal,
		arg.IpAddressExternal,
		arg.Hostname,
		arg.StatusCapacityCpu,
		arg.StatusCapacityMemory,
		arg.StatusCapacityPods,
		arg.StatusAllocatableCpu,
		arg.StatusAllocatableMemory,
		arg.StatusAllocatablePods,
		arg.KubeletVersion,
		arg.NodeConditionsReady,
		arg.NodeConditionsDiskPressure,
		arg.NodeConditionsMemoryPressure,
		arg.NodeConditionsPidPressure,
		arg.NodeConditionsNetworkUnavailable,
	)
	return err
}

const updateService = `-- name: UpdateService :exec
INSERT INTO "Service" (
    "name",
    "namespace",
    "timestamp",
    "labels",
    "creation_timestamp",
    "ports",
    "external_ips",
    "cluster_ip"
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type UpdateServiceParams struct {
	Name              string           `db:"name" json:"name"`
	Namespace         string           `db:"namespace" json:"namespace"`
	Timestamp         pgtype.Timestamp `db:"timestamp" json:"timestamp"`
	Labels            []string         `db:"labels" json:"labels"`
	CreationTimestamp pgtype.Timestamp `db:"creation_timestamp" json:"creationTimestamp"`
	Ports             []string         `db:"ports" json:"ports"`
	ExternalIps       []string         `db:"external_ips" json:"externalIps"`
	ClusterIp         string           `db:"cluster_ip" json:"clusterIp"`
}

func (q *Queries) UpdateService(ctx context.Context, arg UpdateServiceParams) error {
	_, err := q.db.Exec(ctx, updateService,
		arg.Name,
		arg.Namespace,
		arg.Timestamp,
		arg.Labels,
		arg.CreationTimestamp,
		arg.Ports,
		arg.ExternalIps,
		arg.ClusterIp,
	)
	return err
}