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
            ROW_NUMBER() OVER (PARTITION BY name, pod_id ORDER BY timestamp DESC) AS row_number
        FROM
            containers) t
    WHERE
        t.row_number = 1
    ORDER BY name ASC`
  );
  const containers: ContainerList = res.rows;
  return containers;
}

export async function getPodsList(): Promise<PodList> {
  const res = await pool.query(
    `SELECT
        *
    FROM (
        SELECT
            *,
            ROW_NUMBER() OVER (PARTITION BY pod_id ORDER BY timestamp DESC, pod_resource_version DESC) AS row_number
        FROM
            pods) p
    WHERE
        p.row_number = 1
    ORDER BY name ASC`
  );
  const pods: PodList = res.rows;
  return pods;
}

export async function getPodDetails(
  pod_id: string
): Promise<{ pod_data: Pod; containers: ContainerList }> {
  const res = await pool.query(
    "SELECT * FROM pods WHERE pod_id = $1 ORDER BY timestamp DESC LIMIT 1",
    [pod_id]
  );
  const pod_data: Pod = res.rows[0];

  const cont = await pool.query(
    `SELECT
        *
    FROM (
        SELECT
            *,
            ROW_NUMBER() OVER (PARTITION BY name, pod_id ORDER BY timestamp DESC) AS row_number
        FROM
            containers) t
    WHERE
        t.row_number = 1
    AND
        pod_id = $1
    ORDER BY name ASC`,
    [pod_id]
  );
  const containers: ContainerList = cont.rows;

  return { pod_data, containers };
}
