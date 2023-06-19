"use client";
import { Table, Dropdown } from "flowbite-react";
import { ContainerData, ContainerList } from "@/lib/types/ContainerList";
import { Health, HealthIndicatorBadge } from "@/components/health_indicators";
//import { list } from "postcss";
import React, { useState } from "react";

export default function ContainerTable({
  list,
}: {
  list: ContainerList;
}): JSX.Element {
  const [sortedList, setSortedList] = useState([...list]);

  const handleSortAsc = () => {
    const sorted = [...sortedList];
    sorted.sort((a, b) => a.status.localeCompare(b.status));
    setSortedList(sorted);
  };

  const handleSortDsc = () => {
    const sorted = [...sortedList];
    sorted.sort((a, b) => b.status.localeCompare(a.status));
    setSortedList(sorted);
  };
  return (
    <div>
      <Table>
        <Table.Head>
          <Table.HeadCell
            className="bg-green-500 bg-opacity-30 text-left"
            scope="col"
            style={{ width: "30%" }}
          >
            Name
          </Table.HeadCell>
          <Table.HeadCell
            className="bg-green-500 bg-opacity-30 text-left"
            scope="col"
          >
            Image
          </Table.HeadCell>
          <Table.HeadCell
            className="bg-green-500 bg-opacity-30 text-left"
            scope="col"
          >
            <span>
              <Dropdown inline label="STATUS" dismissOnClick={true}>
                <Dropdown.Item>
                  <a onClick={handleSortAsc}>Assending</a>
                </Dropdown.Item>
                <Dropdown.Item>
                  <a onClick={handleSortDsc}>Descending</a>
                </Dropdown.Item>
              </Dropdown>
            </span>
          </Table.HeadCell>
        </Table.Head>
        <Table.Body>
          {sortedList.map((container: ContainerData, index: number) => (
            <Table.Row key={index}>
              <Table.Cell className="whitespace-normal font-medium text-gray-900 dark:text-white !py-2">
                <a
                  href={`/containers/${encodeURIComponent(
                    container.container_id
                  )}`}
                  className="text-decoration-none text-blue-800"
                  id="list"
                >
                  {container.name}
                </a>
              </Table.Cell>
              <Table.Cell className="whitespace-normal font-medium text-gray-900 dark:text-white !py-2">
                {container.image}
              </Table.Cell>
              <Table.Cell className="whitespace-normal font-medium text-gray-900 dark:text-white !py-2">
                <HealthIndicatorBadge status={container.status as Health} />
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </div>
  );
}
