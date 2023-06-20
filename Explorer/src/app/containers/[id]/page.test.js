import { render } from "@testing-library/react";
import Index from "./page";
//eslint-disable-next-line @typescript-eslint/no-unused-vars
import { toBeInTheDocument } from "@testing-library/jest-dom";

jest.mock("../../../lib/db", () => ({
  getContainerDetails: jest.fn((arg) => {
    if (arg != "1") {
      throw "Wrong argument passed to getContainerDetails";
    }
    return Promise.resolve({
      container_event_id: 1,
      container_id: 1,
      timestamp: "2021-08-01 00:00:00",
      pod_id: 1,
      name: "Container 1",
      image: "Image 1",
      status: "Running",
      ports: 1,
    });
  }),
}));

describe("Index", () => {
  it("displays the container", async () => {
    const { getByText } = render(await Index({ params: { id: "1" } }));

    expect(getByText("Container 1")).toBeInTheDocument();
  });
});
