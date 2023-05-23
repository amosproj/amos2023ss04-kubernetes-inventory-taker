CREATE TABLE "cluster" (
  "cluster_event_id" int PRIMARY KEY,
  "cluster_id" int,
  "timestamp" timestamp NOT NULL,
  "name" varchar(255)
);

-- https://kubernetes.io/docs/concepts/architecture/nodes/
CREATE TABLE "nodes" (
  "node_event_id" int PRIMARY KEY,
  "node_id" int,
  "timestamp" timestamp NOT NULL,
  "cluster_id" int,
  "name" varchar(255),
  "ip_address_internal" varchar(255),
  "ip_address_external" varchar(255),
  "status" varchar(50)
);

-- https://kubernetes.io/docs/concepts/workloads/pods/
CREATE TABLE "pods" (
  "pod_event_id" int PRIMARY KEY,
  "pod_id" int,
  "timestamp" timestamp NOT NULL,
  "cluster_id" int,
  "node_id" int,
  "name" varchar(255),
  "status" varchar(50)
);

-- https://kubernetes.io/docs/concepts/containers/
CREATE TABLE "containers" (
  "container_event_id" int PRIMARY KEY,
  "container_id" int,
  "timestamp" timestamp NOT NULL,
  "pod_id" int,
  "name" varchar(255),
  "image" varchar(255),
  "status" varchar(50),
  "ports" varchar(255)
);

-- https://kubernetes.io/docs/concepts/services-networking/service/
CREATE TABLE "services" (
  "service_id" int PRIMARY KEY,
  "cluster_id" int,
  "name" varchar(255),
  "type" varchar(50),
  "port" int
);
