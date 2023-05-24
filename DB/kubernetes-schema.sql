CREATE TABLE "cluster" (
  "cluster_event_id" int PRIMARY KEY,
  "cluster_id" uuid,
  "timestamp" timestamp NOT NULL,
  "name" varchar(255)
);

-- https://kubernetes.io/docs/concepts/architecture/nodes/
CREATE TABLE "node" (
  "node_event_id" int PRIMARY KEY,
  "node_id" uuid,
  "timestamp" timestamp NOT NULL,
  "cluster_id" int,
  "name" varchar(255),
  "ip_address_internal" varchar(255),
  "ip_address_external" varchar(255),
  "status" json
);

-- https://kubernetes.io/docs/concepts/workloads/pods/
CREATE TABLE "pod" (
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
CREATE TABLE "container" (
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
CREATE TABLE "service" (
  "service_id" int PRIMARY KEY,
  "cluster_id" uuid,
  "name" varchar(255),
  "type" varchar(50),
  "port" int
);

-- ALTER TABLE "Service" ADD FOREIGN KEY ("cluster_id") REFERENCES "Cluster" ("cluster_id");
