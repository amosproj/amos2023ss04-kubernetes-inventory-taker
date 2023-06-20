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
};
export type ContainerList = Container[];
