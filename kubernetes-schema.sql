CREATE TABLE "Cluster" (
  "cluster_event_id" int PRIMARY KEY,
  "cluster_id" int,
  "timestamp" datetime NOT NULL;
  "name" varchar(255)
);

-- https://kubernetes.io/docs/concepts/architecture/nodes/
CREATE TABLE "Node" (
  "node_event_id" int PRIMARY KEY,
  "node_id" int,
  "timestamp" datetime NOT NULL;
  "cluster_id" int,
  "name" varchar(255),
  "ip_address_internal" varchar(255),
  "ip_address_external" varchar(255),
  "status" varchar(50)
);

-- https://kubernetes.io/docs/concepts/workloads/pods/
CREATE TABLE "Pod" (
  "pod_event_id" int PRIMARY KEY,
  "pod_id" int,
  "timestamp" datetime NOT NULL;
  "cluster_id" int,
  "node_id" int,
  "name" varchar(255),
  "status" varchar(50)
);

-- https://kubernetes.io/docs/concepts/containers/
CREATE TABLE "Container" (
  "container_event_id" int PRIMARY KEY,
  "container_id" int,
  "timestamp" datetime NOT NULL;
  "pod_id" int,
  "name" varchar(255),
  "image" varchar(255),
  "status" varchar(50),
  "ports" varchar(255),
);

-- https://kubernetes.io/docs/concepts/services-networking/service/
CREATE TABLE "Service" (
  "service_id" int PRIMARY KEY,
  "cluster_id" int,
  "name" varchar(255),
  "type" varchar(50),
  "port" int
);

ALTER TABLE "Service" ADD FOREIGN KEY ("cluster_id") REFERENCES "Cluster" ("cluster_id");
