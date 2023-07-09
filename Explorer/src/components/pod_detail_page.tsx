"use client";

import { Table } from "flowbite-react";
import { H1 } from "@/components/style_elements";
import { HealthIndicatorBadge } from "@/components/health_indicators";
import { Pod } from "@/lib/types/Pod";

export default function PodDetailPage({
  pod_details,
}: {
  pod_details: Pod;
}): JSX.Element {
  return (
    <div>
      <div className="flex">
        <H1 content={"Pod ID " + pod_details.id} />
        <HealthIndicatorBadge status={pod_details.status_phase} />
      </div>
      <div className="flex">
        <div className="w-1/4 w-max">
          <PodDetailsWidget pod_data={pod_details} />
        </div>
      </div>
    </div>
  );
}

function PodDetailsWidget({ pod_data }: { pod_data: Pod }): JSX.Element {
  return (
    <div className="p-0 w-max">
      <h2 className="mt-2 mb-3 text-2xl font-bold">Details</h2>
      <Table>
        <Table.Head>
          <Table.HeadCell className="!py-2 bg-gray-50 dark:bg-gray-800">
            Field
          </Table.HeadCell>
          <Table.HeadCell className="!py-2 bg-gray-30 dark:bg-gray-600">
            Content
          </Table.HeadCell>
        </Table.Head>
        <Table.Body className="divide-y">
          {Object.entries(pod_data).map(([name, value], index) => {
            if (value instanceof Date) {
              value = value.toUTCString();
            } else if (typeof value === "boolean") {
              value = value ? "true" : "false";
            }
            return (
              <Table.Row
                key={index}
                className="bg-white dark:border-gray-700 dark:bg-gray-800"
              >
                {name === "container_status" ? (
                  <>
                    <Table.Cell className="!py-1 whitespace-nowrap font-medium bg-gray-50 dark:bg-gray-800">
                      {name.toUpperCase()}
                    </Table.Cell>
                    <Table.Cell className="whitespace-nowrap font-medium bg-gray-30 dark:bg-gray-600 !py-1">
                      <HealthIndicatorBadge
                        status={pod_data.container_status}
                      />
                    </Table.Cell>
                  </>
                ) : (
                  <>
                    <Table.Cell className="!py-1 whitespace-nowrap font-medium bg-gray-50 dark:bg-gray-800">
                      {name.toUpperCase()}
                    </Table.Cell>
                    <Table.Cell className="!py-1 whitespace-nowrap font-medium bg-gray-30 dark:bg-gray-600">
                      {value}
                    </Table.Cell>
                  </>
                )}
              </Table.Row>
            );
          })}
        </Table.Body>
      </Table>
    </div>
  );
}
