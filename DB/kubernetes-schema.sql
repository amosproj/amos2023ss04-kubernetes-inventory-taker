CREATE TABLE "Clusters"(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "timestamp" timestamp NOT NULL,
  "cluster_id" int,
  "name" text
);

-- https://kubernetes.io/docs/concepts/architecture/nodes/
CREATE TABLE "Nodes"(
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

-- https://kubernetes.io/docs/concepts/workloads/pods/
CREATE TABLE "Pods"(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "timestamp" timestamp NOT NULL,
  "name" text,
  "pod_resource_version" text NOT NULL,
  "pod_id" uuid NOT NULL,
  "node_name" text,
  "namespace" text,
  "status_phase" text,
  "data" json NOT NULL
);

-- https://kubernetes.io/docs/concepts/containers/
CREATE TABLE "Containers"(
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "timestamp" timestamp NOT NULL,
  "container_id" text,
  "pod_id" uuid,
  "name" text,
  "image" text,
  "status" text,
  "ports" text
);

-- https://kubernetes.io/docs/concepts/services-networking/service/
CREATE TABLE "Services"(
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
