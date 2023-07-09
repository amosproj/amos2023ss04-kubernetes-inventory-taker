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
  image_id: string;
  ready: boolean;
  restart_count: number;
  started: boolean;
};
export type ContainerDetails = {
  container: Container;
  status: ContainerStates;
};
export type ContainerList = Container[];
