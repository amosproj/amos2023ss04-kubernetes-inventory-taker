import { Status } from "@/lib/types/Status";

export type Pod = {
  id: number;
  name: string;
  namespace: string;
  status_phase: Status;
};

export type PodList = Pod[];
