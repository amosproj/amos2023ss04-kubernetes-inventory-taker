import "server-only";
import { ContainerDetails, ContainerIndex } from "./types/ContainerDetails";
import { ContainerList, ContainerData } from "./types/ContainerList";
import { Pool } from "pg";

// Example for Container Data Structure to generate page
const container_data = [
  { field: "ID", content: "235480394" },
  { field: "Name", content: "Company Mail" },
  { field: "Status", content: "Error" },
  { field: "Image", content: "postfix:v1.0.3" },
  { field: "Service", content: "Mailserver" },
  { field: "Cluster", content: "DMZ-Cluster" },
  { field: "Node", content: "DMZ-Node-2" },
  { field: "Pod", content: "DMZ-Pod-Mail" },
  { field: "Ports", content: "25:25, 110:110" },
  { field: "Volumes", content: "Mail-Storage-1, Mail-Storage-2" },
  { field: "Age", content: "124 Days 10 Hours 22 Minutes 12 Seconds" },
  { field: "CPU Usage", content: "22 %" },
  { field: "Space Usage", content: "212 GB" },
  { field: "Created on", content: "12.03.2022 12:12:56" },
  { field: "Restart Option", content: "Always" },
];

// Example for Changelog Data Structure to generate page
const changelog_data = [
  {
    status: "running",
    name: "nginx-earthlt",
    port: "8000",
    started: "15 minutes ago",
  },
  {
    status: "running",
    name: "kind bouman",
    port: "3000",
    started: "10 minutes ago",
  },
  { status: "disabled", name: "radis-stack", port: "8000", started: "" },
];

const pool = new Pool({
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  host: process.env.DB_HOST,
  port: parseInt(process.env.DB_PORT || "5432"),
  database: process.env.PGSQL_DATABASE,
});

export async function getContainerDetails(
  container_id: string | undefined
): Promise<ContainerDetails> {
  const adjusted_container_data = structuredClone(container_data);
  if (container_id) {
    const res = (
      await pool.query(
        "SELECT * FROM containers c WHERE container_id = $1 order by timestamp DESC limit 1",
        [container_id]
      )
    ).rows[0];
    adjusted_container_data[ContainerIndex.ID].content = res.container_id;
    adjusted_container_data[ContainerIndex.NAME].content = res.name;
    adjusted_container_data[ContainerIndex.STATUS].content = res.status;
    adjusted_container_data[ContainerIndex.IMAGE].content = res.image;
    adjusted_container_data[ContainerIndex.POD].content = res.pod_id;
  }

  return { fields: adjusted_container_data, changelog: changelog_data };
}

export async function getContainerList(): Promise<ContainerList> {
  const res = await pool.query(
    "SELECT * FROM containers order by timestamp DESC"
  );
  const containers: ContainerData[] = res.rows;
  return containers;
}
