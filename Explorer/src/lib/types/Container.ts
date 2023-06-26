import { Status } from "./Status";

export type Container = {
  id: number;
  timestamp: Date;
  container_id: number;
  pod_id: number;
  name: string;
  image: string;
  status: Status;
  ports: number;
  image_id: string;
  ready: boolean;
  restart_count: number;
  started: boolean;
  state_id: number;
  last_state_id: number;
};
export type ContainerList = Container[];
