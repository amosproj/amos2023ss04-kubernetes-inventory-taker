import "server-only";
import { ContainerDetails, ContainerList } from "./types/Container";
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
  return (
    await pool.query(
      "SELECT * FROM containers c LEFT JOIN container_states cs ON cs.id  = c.state_id WHERE c.container_id = $1 order by timestamp DESC limit 1",
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
