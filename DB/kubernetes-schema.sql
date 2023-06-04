CREATE TABLE "clusters" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "resource_version" text NOT NULL,
  "timestamp" timestamp NOT NULL,
  "cluster_uid" uuid NOT NULL,
  "name" text NOT NULL,
  "data" json
);
-- https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#node-v1-core
-- https://kubernetes.io/docs/concepts/architecture/nodes/
CREATE TABLE "nodes" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "resource_version" text NOT NULL,
  "timestamp" timestamp NOT NULL,
  "node_uid" uuid NOT NULL,
  "name" text,
  "namespace" text,
  "ip_address_internal" text ARRAY,
  "ip_address_external" text ARRAY,
  "hostname" text,
  "status" text,
  "data" json
);
-- https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#pod-v1-core
-- https://kubernetes.io/docs/concepts/workloads/pods/
CREATE TABLE "pods" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "resource_version" text NOT NULL,
  "timestamp" timestamp NOT NULL,
  "pod_uid" uuid NOT NULL,
  "name" text NOT NULL,
  "namespace" text,
  "labels" text[],
  "node_name" text,
  "status" text,
  "data" json
);
-- https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#container-v1-core
-- https://kubernetes.io/docs/concepts/containers/
CREATE TABLE "containers" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  /* "resource_version" string NOT NULL, */
  /* Containers do not seem to have any resource version/metav1 field*/
  "timestamp" timestamp NOT NULL,
  "container_uid" uuid NOT NULL,
  "name" text NOT NULL,
  "pod_uid" uuid,
  "image" text,
  "ports" text,
  "status" text,
  "data" json
);
-- https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#service-v1-core
-- https://kubernetes.io/docs/concepts/services-networking/service/
CREATE TABLE "services" (
  "id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "resource_version" text NOT NULL,
  "timestamp" timestamp NOT NULL,
  "service_uid" uuid NOT NULL,
  "name" text NOT NULL,
  "namespace" text,
  "labels" text[],
  "external_name" text[],
  "status" text,
  "data" json
);
