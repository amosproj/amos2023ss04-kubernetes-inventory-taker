describe("ContainerTable", () => {
  beforeEach(() => {
    // Visit container list page
    cy.visit("/pods");

    // Collect all container links
    cy.get('a[href*="/pods/"]').then(($links) => {
      // Choose a random index
      const randomIndex = Math.floor(Math.random() * $links.length);
      // Get the link at the random index
      const randomLink = $links[randomIndex];
      // Click the random link to navigate to container detail page
      cy.wrap(randomLink).click();
    });
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
