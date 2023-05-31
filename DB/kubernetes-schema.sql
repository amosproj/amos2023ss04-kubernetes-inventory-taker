CREATE TABLE "Cluster" (
  "cluster_event_id" int PRIMARY KEY,
  "cluster_id" int,
  "timestamp" timestamp NOT NULL,
  "name" text
);

-- https://kubernetes.io/docs/concepts/architecture/nodes/
CREATE TABLE "Node" (
  "node_event_id" int PRIMARY KEY,
  "node_id" uuid,
  "timestamp" timestamp NOT NULL,
  "creation_time" timestamp,
  "name" text,
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
  "node_conditions_network_unavailable" text
);

-- https://kubernetes.io/docs/concepts/workloads/pods/
CREATE TABLE "Pod" (
  "pod_resource_version" int NOT NULL,
  "pod_id" uuid NOT NULL,
  "timestamp" timestamp NOT NULL,
  "cluster_id" int,
  "node_id" varchar(50),
  "name" varchar(255),
  "namespace" varchar(50),
  "status" varchar(50),
  "stamp2" timestamp,
  PRIMARY KEY ("pod_resource_version", "pod_id")
);

-- https://kubernetes.io/docs/concepts/containers/
CREATE TABLE "Container" (
  "container_event_id" int PRIMARY KEY,
  "container_id" int,
  "timestamp" timestamp NOT NULL,
  "pod_id" uuid,
  "name" varchar(255),
  "image" varchar(255),
  "status" text,
  "ports" varchar(255)
);

-- https://kubernetes.io/docs/concepts/services-networking/service/
CREATE TABLE "Service" (
  "name" TEXT NOT NULL,
  "namespace" TEXT NOT NULL,
  "timestamp" TIMESTAMP NOT NULL,
  "labels" TEXT[] NOT NULL,
  "creation_timestamp" TIMESTAMP NOT NULL,
  "ports" TEXT ARRAY NOT NULL,
  "external_ips" TEXT[],
  "cluster_ip" TEXT NOT NULL,
  PRIMARY KEY ("name", "namespace")
);

-- ALTER TABLE "Service" ADD FOREIGN KEY ("cluster_id") REFERENCES "Cluster" ("cluster_id");
