import { render, fireEvent } from "@testing-library/react";
import Index from "./page";
//eslint-disable-next-line @typescript-eslint/no-unused-vars
import { toBeInTheDocument } from "@testing-library/jest-dom";

jest.mock("../../lib/db", () => ({
  getPodsList: jest.fn().mockResolvedValue([
    {
      name: "coredns-787d4945fb-2z6lt",
      namespace: "kube-system",
    },
    {
      name: "hello-k8s-2-75fffcd4c4-9pvs6",
      namespace: "default",
    },
  ]),
}));

describe("Index", () => {
  it("displays the pods list", async () => {
    const { getByText } = render(await Index());

    expect(getByText("hello-k8s-2-75fffcd4c4-9pvs6")).toBeInTheDocument();
  });

  it("filters containers based on search term", async () => {
    const { getByPlaceholderText, getAllByRole, getByText } = render(
      await Index()
    );

    const searchInput = getByPlaceholderText("Search...");
    fireEvent.change(searchInput, {
      target: { value: "coredns-787d4945fb-2z6lt" },
    });
    fireEvent.click(getByText("Search"));

    const tableRows = getAllByRole("row");
    expect(tableRows).toHaveLength(2); // Header row + 1 matching row

    fireEvent.change(searchInput, { target: { value: "default" } });
    fireEvent.click(getByText("Search"));

    const updatedTableRows = getAllByRole("row");
    expect(updatedTableRows).toHaveLength(2); // All pods match the search term
  });
});
