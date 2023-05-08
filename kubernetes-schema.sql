CREATE TABLE "Cluster" (
  "cluster_id" int PRIMARY KEY,
  "name" varchar(255)
);

CREATE TABLE "Node" (
  "node_id" int PRIMARY KEY,
  "cluster_id" int,
  "name" varchar(255),
  "ip_address_internal" varchar(255),
  "ip_address_external" varchar(255),
  "status" varchar(50)
);

CREATE TABLE "Pod" (
  "pod_id" int PRIMARY KEY,
  "cluster_id" int,
  "node_id" int,
  "name" varchar(255),
  "status" varchar(50)
);

CREATE TABLE "Container" (
  "container_id" int PRIMARY KEY,
  "pod_id" int,
  "name" varchar(255),
  "image" varchar(255),
  "status" varchar(50)
);

CREATE TABLE "Service" (
  "service_id" int PRIMARY KEY,
  "cluster_id" int,
  "name" varchar(255),
  "type" varchar(50),
  "port" int
);