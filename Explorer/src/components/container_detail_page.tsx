"use client";

import { Table } from "flowbite-react";
import { H1 } from "@/components/style_elements";
import { HealthIndicatorBadge } from "@/components/health_indicators";

// Magic Number Definitions for Container Data Structure
const CONTAINER_ID = 0;
const CONTAINER_NAME = 1;
const CONTAINER_STATUS = 2;
const CONTAINER_IMAGE = 3;
const CONTAINER_SERVICE = 4;
const CONTAINER_CLUSTER = 5;
const CONTAINER_NODE = 6;
const CONTAINER_POD = 7;
const CONTAINER_PORTS = 8;
const CONTAINER_VOLUMES = 9;
const CONTAINER_AGE = 10;
const CONTAINER_CPU_USAGE = 11;
const CONTAINER_SPACE_USAGE = 12;
const CONTAINER_CREATED_ON = 13;
const CONTAINER_RESTART_OPTIONES = 14;

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

export default function ContainerDetailPage(): JSX.Element {
  return (
    <div>
      <div className="flex">
        <H1 content={"Container ID " + container_data[CONTAINER_ID].content} />
        <HealthIndicatorBadge
          status={container_data[CONTAINER_STATUS].content}
        />
      </div>
      <div className="flex">
        <div className="w-1/4 w-max">
          <ContainerDetailsWidget />
        </div>
        <div className="w-1/2 w-max px-8">
          <ContainerWorkLoad />
          <ContainerChangelogWidget />
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

function ContainerChangelogWidget(): JSX.Element {
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

function ContainerDetailsWidget(): JSX.Element {
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
