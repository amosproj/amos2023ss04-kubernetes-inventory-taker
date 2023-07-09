import "server-only";
import { Container, ContainerList } from "./types/Container";
import { Pool } from "pg";
import { Pod, PodList } from "./types/Pod";

const pool = new Pool({
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  host: process.env.DB_HOST,
  port: parseInt(process.env.DB_PORT || "5432"),
  database: process.env.PGSQL_DATABASE,
});

export async function getContainerDetails(
  container_id: string
): Promise<Container | undefined> {
  return (
    await pool.query(
      "SELECT * FROM containers c WHERE container_id = $1 order by timestamp DESC limit 1",
      [container_id]
    )
  ).rows[0];
}

export async function getContainerList(): Promise<ContainerList> {
  const res = await pool.query(
    // returns entry with the last timestamp for every unique (name, pod_id)
    `SELECT
        *
    FROM (
        SELECT
            *,
            RANK() OVER (PARTITION BY name, pod_id ORDER BY timestamp DESC) AS rank
        FROM
            containers) t
    WHERE
        t.rank = 1`
  );
  const containers: ContainerList = res.rows;
  return containers;
}

export async function getPodsList(): Promise<PodList> {
  const res = await pool.query(
    `SELECT
        *
    FROM pods`
  );

  const pods: PodList = res.rows;
  return pods;
}

export async function getPodDetails(id: string): Promise<Pod | undefined> {
  return (
    await pool.query(
      "SELECT p.id, p.timestamp, p.name, p.pod_resource_version, p.pod_id, p.node_name, p.namespace, p.status_phase, p.host_ip, p.pod_ip, p.pod_ips, p.start_time, p.qos_class, c.container_id, c.image, c.status as container_status, c.ports, c.image_id FROM pods p JOIN containers c ON p.pod_id = c.pod_id WHERE p.id = $1 ORDER BY p.timestamp DESC LIMIT 1",
      [id]
    )
  ).rows[0];
}
