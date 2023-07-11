"use client";
import Link from "next/link";
import { Table, Dropdown } from "flowbite-react";
import { Container, ContainerList } from "@/lib/types/Container";
import { HealthIndicatorBadge } from "@/components/health_indicators";
//import { list } from "postcss";
import React, { useState } from "react";

export default function ContainerTable({
  list,
}: {
  list: ContainerList;
}): JSX.Element {
  const [searchTerm, setSearchTerm] = useState<string>("");
  const [sortDirection, setSortDirection] = useState<"asc" | "desc" | "none">(
    "none"
  );
  const [searchActive, setSearchActive] = useState<boolean>(false);
  const sortFn = (a: Container, b: Container) => {
    if (sortDirection === "none") {
      // don't change anything, see https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/sort#description
      return 0;
    }
    if (sortDirection === "asc") {
      return a.status.localeCompare(b.status);
    }
    // must be desc
    return b.status.localeCompare(a.status);
  };
  const displayList = list
    .filter((container) => {
      if (searchTerm === "" || !searchActive) {
        return true;
      }
      return (
        container.image.toLowerCase().includes(searchTerm.toLowerCase()) ||
        container.name.toLowerCase().includes(searchTerm.toLowerCase())
      );
    })
    .sort(sortFn);
  return (
    <div>
      <div className="mb-4 text-right">
        <input
          type="text"
          placeholder="Search..."
          value={searchTerm}
          onChange={(e) => {
            setSearchTerm(e.target.value);
            setSearchActive(false);
          }}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              setSearchActive(true);
            }
          }}
          className="border border-gray-300 px-4 py-2 rounded-md"
        />
        <button
          type="button"
          onClick={() => setSearchActive(true)}
          className="ml-2 px-4 py-2 bg-blue-500 text-white rounded-md"
        >
          Search
        </button>
      </div>
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
              <Dropdown inline label="STATUS">
                <Dropdown.Item onClick={() => setSortDirection("asc")}>
                  Ascending
                </Dropdown.Item>
                <Dropdown.Item onClick={() => setSortDirection("desc")}>
                  Descending
                </Dropdown.Item>
              </Dropdown>
            </span>
          </Table.HeadCell>
        </Table.Head>
        <Table.Body>
          {displayList.map((container: Container, index: number) => (
            <Table.Row key={index}>
              <Table.Cell className="whitespace-normal font-medium text-gray-900 dark:text-white !py-2">
                <Link
                  href={`/containers/${encodeURIComponent(
                    container.container_id
                  )}`}
                  className="text-decoration-none text-blue-800"
                  id="list"
                >
                  {container.name}
                </Link>
              </Table.Cell>
              <Table.Cell className="whitespace-normal font-medium text-gray-900 dark:text-white !py-2">
                {container.image}
              </Table.Cell>
              <Table.Cell className="whitespace-normal font-medium text-gray-900 dark:text-white !py-2">
                <HealthIndicatorBadge status={container.status} />
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </div>
  );
}
