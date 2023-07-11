"use client";
import Link from "next/link";
import { Table, Dropdown } from "flowbite-react";
import { Pod, PodList } from "@/lib/types/Pod";
import React, { useState } from "react";

export default function PodTable({ list }: { list: PodList }): JSX.Element {
  const [searchTerm, setSearchTerm] = useState<string>("");
  const [sortDirection, setSortDirection] = useState<"asc" | "desc" | "none">(
    "none"
  );
  const [searchActive, setSearchActive] = useState<boolean>(false);
  const sortFn = (a: Pod, b: Pod) => {
    if (sortDirection === "none") {
      // don't change anything, see https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/sort#description
      return 0;
    }
    if (sortDirection === "asc") {
      return a.status_phase.localeCompare(b.status_phase);
    }
    // must be desc
    return b.status_phase.localeCompare(a.status_phase);
  };
  const displayList = list
    .filter((container) => {
      if (searchTerm === "" || !searchActive) {
        return true;
      }
      return (
        container.namespace.toLowerCase().includes(searchTerm.toLowerCase()) ||
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
            Namespace
          </Table.HeadCell>
          <Table.HeadCell
            className="bg-green-500 bg-opacity-30 text-left"
            scope="col"
          >
            <span>
              <Dropdown inline label="STATUS" dismissOnClick={true}>
                <Dropdown.Item>
                  <button onClick={() => setSortDirection("asc")}>
                    Ascending
                  </button>
                </Dropdown.Item>
                <Dropdown.Item>
                  <button onClick={() => setSortDirection("desc")}>
                    Descending
                  </button>
                </Dropdown.Item>
              </Dropdown>
            </span>
          </Table.HeadCell>
        </Table.Head>
        <Table.Body>
          {displayList.map((pod: Pod, index: number) => (
            <Table.Row key={index}>
              <Table.Cell className="whitespace-normal font-medium text-gray-900 dark:text-white !py-2">
                <Link
                  href={`/pods/${encodeURIComponent(pod.pod_id)}`}
                  className="text-decoration-none text-blue-800"
                  id="list"
                >
                  {pod.name}
                </Link>
              </Table.Cell>
              <Table.Cell className="whitespace-normal font-medium text-gray-900 dark:text-white !py-2">
                {pod.namespace}
              </Table.Cell>
              <Table.Cell className="whitespace-normal font-medium text-gray-900 dark:text-white !py-2">
                {/* <HealthIndicatorBadge status={pod.status} /> */}
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </div>
  );
}
