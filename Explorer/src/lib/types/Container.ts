import { Status } from "./Status";
export type ContainerStates =
  | { kind: "running"; started_at: Date }
  | {
      kind: "terminated";
      container_id: number;
      reason: string;
      message: string;
      started_at: Date;
      finished_at: Date;
      exit_code: number;
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
  image_id: string;
  ready: boolean;
  restart_count: number;
  started: boolean;
  current_state: ContainerStates;
  last_fail_state: ContainerStates | undefined;
};
export type ContainerList = Container[];
