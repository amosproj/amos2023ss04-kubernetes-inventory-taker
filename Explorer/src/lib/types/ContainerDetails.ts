// Magic Number Definitions for Container Data Structure
export enum ContainerIndex {
  ID,
  NAME,
  STATUS,
  IMAGE,
  SERVICE,
  CLUSTER,
  NODE,
  POD,
  PORTS,
  VOLUMES,
  AGE,
  CPU_USAGE,
  SPACE_USAGE,
  CREATED_ON,
  RESTART_OPTIONS,
}

export type ContainerDetails = {
  fields: Array<{ field: string; content: string }>;
  changelog: Array<{
    status: string;
    name: string;
    port: string;
    started: string;
  }>;
};
