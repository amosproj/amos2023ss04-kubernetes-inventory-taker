"use client";
import Link from "next/link";
import { Table, Dropdown } from "flowbite-react";
import { Pod, PodList } from "@/lib/types/Pod";
import React, { useState } from "react";
import { HealthIndicatorBadge } from "./health_indicators";

export default function PodTable({ list }: { list: PodList }): JSX.Element {
  const [searchTerm, setSearchTerm] = useState("");
  const [filteredPods, setFilteredPods] = useState<PodList>(list);

  const handleSearch = () => {
    const filtered = list.filter(
      (pod) =>
        pod.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        pod.namespace.toLowerCase().includes(searchTerm.toLowerCase())
    );
    setFilteredPods(filtered);
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(e.target.value);
  };

  const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      handleSearch();
    }
  };

  //   const handleSortAsc = () => {
  //     const sorted = [...filteredPods];
  //     sorted.sort((a, b) => a.status.localeCompare(b.status));
  //     setFilteredPods(sorted);
  //   };

  //   const handleSortDsc = () => {
  //     const sorted = [...filteredPods];
  //     sorted.sort((a, b) => b.status.localeCompare(a.status));
  //     setFilteredPods(sorted);
  //   };

  return (
    <div>
      <div className="mb-4 text-right">
        <input
          type="text"
          placeholder="Search..."
          value={searchTerm}
          onChange={handleInputChange}
          onKeyPress={handleKeyPress}
          className="border border-gray-300 px-4 py-2 rounded-md"
        />
        <button
          type="button"
          onClick={handleSearch}
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
                {/* <Dropdown.Item>
                  <button onClick={() => handleSortAsc()}>Ascending</button>
                </Dropdown.Item>
                <Dropdown.Item>
                  <button onClick={() => handleSortDsc()}>Descending</button>
                </Dropdown.Item> */}
              </Dropdown>
            </span>
          </Table.HeadCell>
        </Table.Head>
        <Table.Body>
          {filteredPods.map((pod: Pod, index: number) => (
            <Table.Row key={index}>
              <Table.Cell className="whitespace-normal font-medium text-gray-900 dark:text-white !py-2">
                <Link
                  href={`/pods/${encodeURIComponent(pod.id)}`}
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
                <HealthIndicatorBadge status={pod.status_phase} />
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </div>
  );
}
