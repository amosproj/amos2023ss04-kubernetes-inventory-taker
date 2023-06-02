import "server-only";
import { Client } from "pg";
import { ContainerDetails } from "./types/ContainerDetails";

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

export async function getData(): Promise<ContainerDetails> {
  const client = new Client();
  await client.connect();

  const res = await client.query("SELECT $1::text as message", [
    "Hello world!",
  ]);
  await client.end();
  return { fields: container_data, changelog: changelog_data };
}
