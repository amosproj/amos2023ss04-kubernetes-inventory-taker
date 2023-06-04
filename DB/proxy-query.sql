-- name: UpdateService :exec
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
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
-- -- name: UpdatePod :one
-- INSERT INTO "Pod" (
--     "pod_resource_version", "pod_id", "timestamp", "cluster_id", "node_id", "name", "namespace","status", "stamp2"
-- ) VALUES (
--     $1, $2, $3, $4, $5, $6, $7, $8, $9
-- )
-- RETURNING "pod_resource_version", "pod_id", "timestamp", "cluster_id", "node_id", "name", "namespace","status", "stamp2";
--
-- -- name: UpdateContainer :one
-- INSERT INTO "Container" (
--     "container_event_id", "container_id", "timestamp", "pod_id", "name", "image", "status", "ports"
-- ) VALUES (
--     $1, $2, $3, $4, $5, $6, $7, $8
-- )
-- RETURNING "container_event_id", "container_id", "timestamp", "pod_id", "name", "image", "status", "ports";
-- name: UpdateNode :exec
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
  );
-- name: UpdateCluster :exec
INSERT INTO "Cluster" (
    "cluster_id",
    "timestamp",
    "name"
  )
VALUES ($1, $2, $3);