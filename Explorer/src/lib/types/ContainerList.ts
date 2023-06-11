export type ContainerData = {
  container_event_id: number;
  container_id: number;
  timestamp: string;
  pod_id: number;
  name: string;
  image: string;
  status: string;
  ports: number;
};

export type ContainerList = ContainerData[];
