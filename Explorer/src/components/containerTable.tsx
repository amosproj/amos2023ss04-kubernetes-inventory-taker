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
  const [searchTerm, setSearchTerm] = useState("");
  const [filteredContainers, setFilteredContainers] =
    useState<ContainerList>(list);

  const handleSearch = () => {
    const filtered = list.filter(
      (container) =>
        container.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        container.image.toLowerCase().includes(searchTerm.toLowerCase())
    );
    setFilteredContainers(filtered);
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(e.target.value);
  };

  const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      handleSearch();
    }
  };

  const handleSortAsc = () => {
    const sorted = [...filteredContainers];
    sorted.sort((a, b) => a.status.localeCompare(b.status));
    setFilteredContainers(sorted);
  };

  const handleSortDsc = () => {
    const sorted = [...filteredContainers];
    sorted.sort((a, b) => b.status.localeCompare(a.status));
    setFilteredContainers(sorted);
  };
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
            Image
          </Table.HeadCell>
          <Table.HeadCell
            className="bg-green-500 bg-opacity-30 text-left"
            scope="col"
          >
            <span>
              <Dropdown inline label="STATUS">
                <Dropdown.Item onClick={() => handleSortAsc()}>
                  Ascending
                </Dropdown.Item>
                <Dropdown.Item onClick={() => handleSortDsc()}>
                  Descending
                </Dropdown.Item>
              </Dropdown>
            </span>
          </Table.HeadCell>
        </Table.Head>
        <Table.Body>
          {filteredContainers.map((container: Container, index: number) => (
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
