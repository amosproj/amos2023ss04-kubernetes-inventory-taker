CREATE TABLE "Cluster" (
  "cluster_id" int PRIMARY KEY,
  "name" varchar(255)
);

-- https://kubernetes.io/docs/concepts/architecture/nodes/
CREATE TABLE "Node" (
  "node_id" int PRIMARY KEY,
  "cluster_id" int,
  "name" varchar(255),
  "ip_address_internal" varchar(255),
  "ip_address_external" varchar(255),
  "status" varchar(50)
);

-- https://kubernetes.io/docs/concepts/workloads/pods/
CREATE TABLE "Pod" (
  "pod_id" int PRIMARY KEY,
  "cluster_id" int,
  "node_id" int,
  "name" varchar(255),
  "status" varchar(50)
);

-- https://kubernetes.io/docs/concepts/containers/
CREATE TABLE "Container" (
  "container_id" int PRIMARY KEY,
  "pod_id" int,
  "name" varchar(255),
  "image" varchar(255),
  "status" varchar(50)
);

-- https://kubernetes.io/docs/concepts/services-networking/service/
CREATE TABLE "Service" (
  "service_id" int PRIMARY KEY,
  "cluster_id" int,
  "name" varchar(255),
  "type" varchar(50),
  "port" int
);

ALTER TABLE "Node" ADD FOREIGN KEY ("cluster_id") REFERENCES "Cluster" ("cluster_id");

ALTER TABLE "Pod" ADD FOREIGN KEY ("cluster_id") REFERENCES "Cluster" ("cluster_id");

ALTER TABLE "Pod" ADD FOREIGN KEY ("node_id") REFERENCES "Node" ("node_id");

ALTER TABLE "Container" ADD FOREIGN KEY ("pod_id") REFERENCES "Pod" ("pod_id");

ALTER TABLE "Service" ADD FOREIGN KEY ("cluster_id") REFERENCES "Cluster" ("cluster_id");
