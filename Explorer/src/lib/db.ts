import "server-only";
import { ContainerDetails, Container, ContainerList } from "./types/Container";
import { Pool } from "pg";

const pool = new Pool({
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  host: process.env.DB_HOST,
  port: parseInt(process.env.DB_PORT || "5432"),
  database: process.env.PGSQL_DATABASE,
});

export async function getContainerDetails(
  container_id: string
): Promise<ContainerDetails | undefined> {
  const res1 = (
    await pool.query(
      "SELECT * FROM containers c WHERE c.container_id = $1 order by timestamp DESC limit 1",
      [container_id]
    )
  ).rows[0];
  if (!res1) {
    // No container found
    return undefined;
  }
  const container: Container = {
    id: res1.id,
    timestamp: res1.timestamp,
    container_id: res1.container_id,
    pod_id: res1.container_id,
    name: res1.name,
    image: res1.image,
    status: res1.status,
    // ports: 0,
    image_id: res1.image_id,
    ready: res1.ready,
    restart_count: res1.restart_count,
    started: res1.started,
    // state_id: 0,
    // last_state_id: 0,
  };

  return { container, status: {} };
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
