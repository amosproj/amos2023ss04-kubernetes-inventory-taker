"use client";

import { Table } from "flowbite-react";
import { H1 } from "@/components/style_elements";
import { Health, HealthIndicatorBadge } from "@/components/health_indicators";
import {
  ChangeLogEntry,
  ContainerDetails,
  ContainerIndex,
} from "@/lib/types/ContainerDetails";

export default function ContainerDetailPage({
  container_details,
}: {
  container_details: ContainerDetails;
}): JSX.Element {
  return (
    <div>
      <div className="flex">
        <H1
          content={
            "Container ID " +
            container_details.fields[ContainerIndex.ID].content
          }
        />
        <HealthIndicatorBadge
          status={
            container_details.fields[ContainerIndex.STATUS].content as Health
          }
        />
      </div>
      <div className="flex">
        <div className="w-1/4 w-max">
          <ContainerDetailsWidget container_data={container_details.fields} />
        </div>
        <div className="w-1/2 w-max px-8">
          <ContainerWorkLoad />
          <ContainerChangelogWidget
            changelog_data={container_details.changelog}
          />
        </div>
      </div>
    </div>
  );
}

function ContainerWorkLoad(): JSX.Element {
  return (
    <div className="flex row">
      <div className="pr-4">
        <header>
          <h3 className="mt-2 mb-3 text-2xl font-bold">CPU Work Load</h3>
        </header>
        <div>
          <a
            href="#"
            className="block max-w-xs p-12 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
          >
            <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              CPU Usage
            </h5>
            <p className="font-normal text-gray-700 text-center dark:text-gray-400">
              15.0%
            </p>
          </a>
        </div>
      </div>

      <div className="pl-4">
        <header>
          <h4 className="mt-2 mb-3 text-2xl font-bold">Memory Work Load</h4>
        </header>
        <div>
          <a
            href="#"
            className="block max-w-xs p-12 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
          >
            <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              Memory Usage
            </h5>
            <p className="font-normal text-gray-700 text-center dark:text-gray-400">
              60.08%
            </p>
          </a>
        </div>
      </div>
    </div>
  );
}

function ContainerChangelogWidget({
  changelog_data,
}: {
  changelog_data: Array<ChangeLogEntry>;
}): JSX.Element {
  return (
    <div className="p-0">
      <section>
        <header>
          <h2 className="mt-2 mb-3 text-2xl font-bold">Changelog</h2>
        </header>
        <div className="">
          <Table>
            <Table.Head>
              <Table.HeadCell className="!py-2 bg-gray-50 dark:bg-gray-800">
                Status
              </Table.HeadCell>
              <Table.HeadCell className="!py-2 bg-gray-30 dark:bg-gray-600">
                Name
              </Table.HeadCell>
              <Table.HeadCell className="!py-2 bg-gray-30 dark:bg-gray-600">
                Port
              </Table.HeadCell>
              <Table.HeadCell className="!py-2 bg-gray-30 dark:bg-gray-600">
                Started
              </Table.HeadCell>
            </Table.Head>
            <Table.Body className="divide-y">
              {changelog_data.map((entry, index) => (
                <Table.Row
                  key={index}
                  className="bg-white dark:border-gray-700 dark:bg-gray-800"
                >
                  <Table.Cell className="!py-1 whitespace-nowrap font-medium bg-gray-50 dark:bg-gray-800">
                    {entry.status == "running" ? (
                      <span className="flex w-3 h-3 bg-green-500 rounded-full"></span>
                    ) : (
                      <span className="flex w-3 h-3 bg-red-500 rounded-full"></span>
                    )}
                  </Table.Cell>
                  <Table.Cell className="!py-1 whitespace-nowrap font-medium bg-gray-30 dark:bg-gray-600">
                    {entry.name}
                  </Table.Cell>
                  <Table.Cell className="!py-1 whitespace-nowrap font-medium bg-gray-30 dark:bg-gray-600">
                    {entry.port}
                  </Table.Cell>
                  <Table.Cell className="!py-1 whitespace-nowrap font-medium bg-gray-30 dark:bg-gray-600">
                    {entry.started}
                  </Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table>
        </div>
      </section>
    </div>
  );
}

function ContainerDetailsWidget({
  container_data,
}: {
  container_data: Array<{ field: string; content: string }>;
}): JSX.Element {
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
          {container_data.map((entry, index) => (
            <Table.Row
              key={index}
              className="bg-white dark:border-gray-700 dark:bg-gray-800"
            >
              <Table.Cell className="!py-1 whitespace-nowrap font-medium bg-gray-50 dark:bg-gray-800">
                {entry.field}
              </Table.Cell>
              <Table.Cell className="!py-1 whitespace-nowrap font-medium bg-gray-30 dark:bg-gray-600">
                {entry.content}
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </div>
  );
}
