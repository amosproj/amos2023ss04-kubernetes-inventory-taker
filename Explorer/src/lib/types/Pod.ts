import { Status } from "./Status";

export type Pod = {
  id: number;
  timestamp: Date;
  name: string;
  pod_resource_version: string;
  pod_id: string;
  node_name: string;
  namespace: string;
  status_phase: Status;
  host_ip: string;
  pod_ip: string;
  pod_ips: string;
  start_time: Date;
  qos_class: string;
  container_status: Status;
};

export type PodList = Pod[];
