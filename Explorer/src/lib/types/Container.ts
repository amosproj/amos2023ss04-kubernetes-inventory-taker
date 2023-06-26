import { Status } from "./Status";
type ContainerStates =
  | { kind: "running"; started_at: Date }
  | {
      kind: "terminated";
      container_id: number;
      exit_code: number;
      finished_at: Date;
      signal: number;
    }
  | { kind: "waiting"; message: string; reason: string };

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
export type ContainerDetails = Container & ContainerStates;
export type ContainerList = Container[];
