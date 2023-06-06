export type ContainerData = {
  container_event_id: string;
  container_id: string;
  timestamp: string;
  pod_id: string;
  name: string;
  image: string;
  status: string;
  ports: string;
};

export type ContainerList = {
  containers: Array<ContainerData>;
};
