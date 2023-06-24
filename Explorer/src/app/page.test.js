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
  ]),
}));

describe("Index", () => {
  it("displays the container list", async () => {
    const { getByText } = render(await Index());

    expect(getByText("Container 1")).toBeInTheDocument();
  });

  it("filters containers based on search term", async () => {
    const { getByPlaceholderText, getAllByRole } = render(await Index());

    const searchInput = getByPlaceholderText("Search...");
    fireEvent.change(searchInput, { target: { value: "Container 1" } });

    const tableRows = getAllByRole("row");
    expect(tableRows).toHaveLength(2); // Header row + 1 matching row

    fireEvent.change(searchInput, { target: { value: "image" } });

    const updatedTableRows = getAllByRole("row");
    expect(updatedTableRows).toHaveLength(2); // All containers match the search term


  });
});
