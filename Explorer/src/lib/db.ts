import "server-only";
import { Container, ContainerList, ContainerStates } from "./types/Container";
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
  const container_row = (
    await pool.query(
      "SELECT * FROM containers c WHERE c.container_id = $1 order by timestamp DESC limit 1",
      [container_id]
    )
  ).rows[0];
  if (!container_row) {
    // No container found
    return undefined;
  }
  const current_state = await getContainerState(container_row.state_id);
  if (!current_state) {
    return undefined;
  }
  const last_fail_state = await getContainerState(container_row.last_state_id);
  const container: Container = {
    id: container_row.id,
    timestamp: container_row.timestamp,
    container_id: container_row.container_id,
    pod_id: container_row.pod_id,
    name: container_row.name,
    image: container_row.image,
    status: container_row.status,
    // ports: 0,
    image_id: container_row.image_id,
    ready: container_row.ready,
    restart_count: container_row.restart_count,
    started: container_row.started,
    // state_id: 0,
    // last_state_id: 0,
    current_state,
    last_fail_state,
  };
  return container;
}

async function getContainerState(
  state_id: number
): Promise<ContainerStates | undefined> {
  const row = (
    await pool.query("SELECT * FROM container_states cs WHERE cs.id = $1", [
      state_id,
    ])
  ).rows[0];
  if (!row) {
    return undefined;
  }

  switch (row.kind) {
    case "Waiting":
      return { kind: "waiting", message: row.message, reason: row.reason };
    case "Terminated":
      return {
        kind: "terminated",
        container_id: row.container_id,
        reason: row.reason,
        message: row.message,
        started_at: row.started_at,
        finished_at: row.finished_at,
        exit_code: row.exit_code,
        signal: row.signal,
      };
    case "Running":
      return { kind: "running", started_at: row.started_at };
  }
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
