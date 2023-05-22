import React, { useState } from "react";
import { Table, Toast } from "flowbite-react";

interface Container {
  name: string;
  image: string;
}

interface StripedTableProps {
  containers: Container[];
}

const ContainerTable: React.FC<StripedTableProps> = ({ containers }) => {
  const [showToast, setShowToast] = useState(false);

  const handleClick = (e: React.MouseEvent) => {
    e.preventDefault();
    setShowToast(true);
  };

  return (
    <>
      {showToast && (
        <div className="py-4">
          <Toast className="bg-orange-500 bg-opacity-50">
            <div className="ml-3 text-sm font-normal">
              your container detail page could be here
            </div>
            <Toast.Toggle onClick={() => setShowToast(false)} />
          </Toast>
        </div>
      )}
      <Table striped>
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
          {containers.map((container, index) => (
            <Table.Row key={index}>
              <Table.Cell className="whitespace-normal font-medium text-gray-900 dark:text-white">
                <a
                  href="#"
                  className="text-decoration-none text-blue-800"
                  onClick={handleClick}
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
    </>
  );
};

export default ContainerTable;
