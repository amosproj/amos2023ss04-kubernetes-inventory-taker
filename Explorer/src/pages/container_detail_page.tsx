"use client";

import { DarkThemeToggle } from "flowbite-react";
import {
  Table,
  Badge,
} from "flowbite-react";

export default function Index(): JSX.Element {
    return (
        <div className="p-6">
            {/*<DarkThemeToggle />*/}
            <ContainerDetailPage />
        </div>
    );
}

function H1({content}: any): JSX.Element {
    return (
        <h1 className="mb-4 my-2 text-3xl font-bold leading-none tracking-tight text-gray-900 dark:text-white">{content}</h1>
    );
}

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
    { field: "Status", content: "Running" },
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

function ContainerDetailPage(): JSX.Element {

    function get_container_status_color() {

        switch (container_data[CONTAINER_STATUS].content) {
            case "Running":
                return "success";
            case "Stopped":
                return "gray";
            case "Error":
                return "failure";
            case "Warning":
                return "warning";
        }

        return "";
    }

    return (
        <div>
            <div className="flex">
                <H1 content={ "Container ID " + container_data[CONTAINER_ID].content } />
                <Badge color={ get_container_status_color() } className="!text-2xl ml-2 mt-1">
                    { container_data[CONTAINER_STATUS].content }
                </Badge>
            </div>
            <div className="flex">
                <div className="w-1/4 w-max">
                    <ContainerDetailsWidget/>
                </div>
                <div className="w-1/4 w-max px-8">
                    <ContainerChangelogWidget />
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
                    <h2 className="mt-2 mb-3 text-2xl font-bold">
                        Changelog
                    </h2>
                </header>
                    <div>
                    </div>
            </section>
        </div>
    );
}


function ContainerDetailsWidget(): JSX.Element {
    return (
        <div className="p-0 w-max">
            <h2 className="mt-2 mb-3 text-2xl font-bold">
                Details
            </h2>
            <Table>
                <Table.Head>
                    <Table.HeadCell className="!py-2 bg-gray-50 dark:bg-gray-800">Field</Table.HeadCell>
                    <Table.HeadCell className="!py-2 bg-gray-30 dark:bg-gray-600">Content</Table.HeadCell>
                </Table.Head>
                <Table.Body className="divide-y">
                    {container_data.map((entry, index) => (
                        <Table.Row className="bg-white dark:border-gray-700 dark:bg-gray-800">
                            <Table.Cell className="!py-1 whitespace-nowrap font-medium bg-gray-50 dark:bg-gray-800">{ entry.field }</Table.Cell>
                            <Table.Cell className="!py-1 whitespace-nowrap font-medium bg-gray-30 dark:bg-gray-600">{ entry.content }</Table.Cell>
                        </Table.Row>
                    ))}
                </Table.Body>
            </Table>
        </div>
    );
}