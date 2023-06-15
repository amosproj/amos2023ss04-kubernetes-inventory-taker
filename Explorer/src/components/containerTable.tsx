"use client";
import { Table } from "flowbite-react";
import { ContainerData, ContainerList } from "@/lib/types/ContainerList";
//import { list } from "postcss";

export default function ContainerTable({
  list,
}: {
  list: ContainerList;
}): JSX.Element {
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
        </Table.Head>
        <Table.Body>
          {list.map((container: ContainerData, index: number) => (
            <Table.Row key={index}>
              <Table.Cell className="whitespace-normal font-medium text-gray-900 dark:text-white">
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
              <Table.Cell className="whitespace-normal font-medium text-gray-900 dark:text-white">
                {container.image}
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </div>
  );
}
