describe("ContainerTable", () => {
  beforeEach(() => {
    cy.visit("/pods/1"); // Replace with the URL of your application
  });

  it("displays the expected fields in the table", () => {
    cy.get("table").should("exist");

    const expectedRows = [
      { field: "FIELD" },
      { field: "ID" },
      { field: "TIMESTAMP" },
      { field: "NAME" },
      { field: "POD_RESOURCE_VERSION" },
      { field: "POD_ID" },
      { field: "NODE_NAME" },
      { field: "NAMESPACE" },
      { field: "STATUS_PHASE" },
      { field: "HOST_IP" },
      { field: "POD_IP" },
      { field: "POD_IPS" },
      { field: "START_TIME" },
      { field: "QOS_CLASS" },
      { field: "CONTAINER_ID" },
      { field: "IMAGE" },
      { field: "CONTAINER_STATUS" },
      { field: "PORTS" },
      { field: "IMAGE_ID" },
    ];

    cy.get("table tr").should("have.length", expectedRows.length);
  });
});
