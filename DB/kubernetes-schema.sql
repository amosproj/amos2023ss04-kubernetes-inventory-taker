CREATE TABLE clusters(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "timestamp" timestamp NOT NULL,
  "cluster_id" int,
  "name" text
);

-- https://kubernetes.io/docs/concepts/services-networking/service/
CREATE TABLE services(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "timestamp" timestamp NOT NULL,
  "name" text NOT NULL,
  "namespace" text NOT NULL,
  "labels" text ARRAY NOT NULL,
  "creation_timestamp" timestamp NOT NULL,
  "ports" text ARRAY NOT NULL,
  "external_ips" text ARRAY,
  "cluster_ip" text NOT NULL,
  "data" json NOT NULL
);

-- https://kubernetes.io/docs/concepts/workloads/pods/
CREATE TABLE pods(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "timestamp" timestamp NOT NULL,
  "name" text,
  "pod_resource_version" text NOT NULL,
  "pod_id" uuid NOT NULL,
  "node_name" text,
  "namespace" text,
  "status_phase" text,
  "data" json NOT NULL,
  "host_ip" text,
  "pod_ip" text,
  "pod_ips" text ARRAY,
  "start_time" timestamp,
  "qos_class" text
);

CREATE TABLE pod_status_conditions(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "pod_id" int REFERENCES pods(id) ON DELETE CASCADE,
  "type" text,
  "status" text,
  "last_probe_time" timestamp,
  "last_transition_time" timestamp,
  "reason" text,
  "message" text
);

CREATE TABLE pod_volumes(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "pod_id" int REFERENCES pods(id) ON DELETE CASCADE,
  "type" text,
  "name" text,
  "persistent_claim_name" text,
  "read_only" bool
);

-- https://kubernetes.io/docs/concepts/architecture/nodes/
CREATE TABLE nodes(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "timestamp" timestamp NOT NULL,
  "name" text,
  "node_id" uuid,
  "creation_time" timestamp,
  "ip_address_internal" text ARRAY,
  "ip_address_external" text ARRAY,
  "hostname" text,
  "status_capacity_cpu" text,
  "status_capacity_memory" text,
  "status_capacity_pods" text,
  "status_allocatable_cpu" text,
  "status_allocatable_memory" text,
  "status_allocatable_pods" text,
  "kubelet_version" text,
  "node_conditions_ready" text,
  "node_conditions_disk_pressure" text,
  "node_conditions_memory_pressure" text,
  "node_conditions_pid_Pressure" text,
  "node_conditions_network_unavailable" text,
  "data" json NOT NULL
);

CREATE TABLE "container_states"(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "kind" text,
  "started_at" time,
  "container_id" text,
  "exit_code" int,
  "finished_at" time,
  "message" text,
  "reason" text,
  "signal" int
);

-- https://kubernetes.io/docs/concepts/containers/
CREATE TABLE containers(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "timestamp" timestamp NOT NULL,
  "container_id" text,
  "pod_id" uuid,
  "name" text,
  "image" text,
  "status" text,
  "ports" text,
  "image_id" text,
  "ready" bool,
  "restart_count" int,
  "started" bool,
  "state_id" int REFERENCES "container_states"(id),
  "last_state_id" int REFERENCES "container_states"(id)
);

CREATE TABLE "volume_devices"(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "container_id" int REFERENCES "containers"("id") NOT NULL,
  "device_path" text,
  "name" text
);

CREATE TABLE "volume_mounts"(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "container_id" int REFERENCES "containers"("id") NOT NULL,
  "mount_path" text,
  "mount_propagation" text,
  "name" text,
  "read_only" bool,
  "sub_path" text,
  "sub_path_expr" text
);

CREATE TABLE "container_ports"(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "container_id" int REFERENCES "containers"("id") NOT NULL,
  "container_port" int,
  "host_ip" text,
  "host_port" int,
  "name" text,
  "protocol" text
);

-- https://kubernetes.io/docs/concepts/storage/persistent-volumes/
CREATE TABLE persistent_volumes(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "timestamp" timestamp NOT NULL,
  "labels" text ARRAY,
  "name" text NOT NULL,
  "namespace" text NOT NULL,
  "uid" text NOT NULL,
  "creation_timestamp" timestamp NOT NULL,
  "deletion_timestamp" timestamp,
  "access_modes" text ARRAY,
  "capacity" text,
  "mount_options" text ARRAY,
  "storage_class_name" text,
  "volume_mode" text,
  "message" text,
  "phase" text,
  "reason" text
);

CREATE TABLE persistent_volume_claims(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "timestamp" timestamp NOT NULL,
  "labels" text ARRAY,
  "name" text NOT NULL,
  "namespace" text NOT NULL,
  "uid" uuid,
  "creation_timestamp" timestamp NOT NULL,
  "deletion_timestamp" timestamp,
  "access_modes" text ARRAY,
  "storage_class_name" text,
  "volume_mode" text,
  "volume_name" text,
  "capacity" text,
  "phase" text,
  "resize_status" text
);

CREATE TABLE persistent_volume_claims_conditions(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "persistent_volume_claim_id" int REFERENCES "persistent_volume_claims" ("id") NOT NULL,
  "last_probe_time" timestamp,
  "last_transition_time" timestamp,
  "message" text,
  "reason" text,
  "status" text,
  "type" text
);
