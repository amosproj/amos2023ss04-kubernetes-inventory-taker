import { render, fireEvent } from "@testing-library/react";
import Index from "./page";
//eslint-disable-next-line @typescript-eslint/no-unused-vars
import { toBeInTheDocument } from "@testing-library/jest-dom";

jest.mock("../lib/db", () => ({
  getContainerList: jest.fn().mockResolvedValue([
    {
      container_event_id: 1,
      container_id: 1,
      timestamp: "2021-08-01 00:00:00",
      pod_id: 1,
      name: "Container 1",
      image: "Image 1",
      status: "Running",
      ports: 1,
    },
    {
      container_event_id: 2,
      container_id: 2,
      timestamp: "2021-08-01 00:00:00",
      pod_id: 2,
      name: "Container 2",
      image: "Image 2",
      status: "Pending",
      ports: 2,
    },

  ]),
}));

describe("Index", () => {
  it("displays the container list", async () => {
    const { getByText } = render(await Index());

    expect(getByText("Container 1")).toBeInTheDocument();
  });

  it("filters containers based on search term", async () => {
    const { getByPlaceholderText, getAllByRole, getByText } = render(await Index());

    const searchInput = getByPlaceholderText("Search...");
    fireEvent.change(searchInput, { target: { value: "Container 1" } });
    fireEvent.click(getByText("Search"));

    const tableRows = getAllByRole("row");
    expect(tableRows).toHaveLength(2); // Header row + 1 matching row

    fireEvent.change(searchInput, { target: { value: "Image" } });
    fireEvent.click(getByText("Search"));

    const updatedTableRows = getAllByRole("row");
    expect(updatedTableRows).toHaveLength(3); // All containers match the search term


  });
});
