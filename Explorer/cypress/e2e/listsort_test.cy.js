describe("ContainerTable", () => {
  beforeEach(() => {
    cy.visit("/"); // Replace with the URL of your application
  });

  it("should sort the list in ascending order", () => {
    cy.contains("button", "STATUS").click({ force: true });
    cy.contains("button", "Ascending").click({ force: true });
    cy.get("td:nth-child(3)") // Assuming the status column is the third column
      .invoke("text")
      .then((statuses) => {
        const sortOrder = ["Running", "Waiting", "Terminated", "Error"];
        const sortedStatuses = [...statuses]
          .sort((a, b) => sortOrder.indexOf(a) - sortOrder.indexOf(b))
          .join("");
        expect(sortedStatuses).to.equal(statuses);
      });
  });

  it("should sort the list in descending order", () => {
    cy.contains("button", "STATUS").click({ force: true });
    cy.contains("button", "Descending").click({ force: true });
    cy.get("td:nth-child(3)") // Assuming the status column is the third column
      .invoke("text")
      .then((statuses) => {
        const sortOrder = ["Running", "Waiting", "Terminated", "Error"];
        const sortedStatuses = [...statuses]
          .sort((a, b) => sortOrder.indexOf(b) - sortOrder.indexOf(a))
          .join("");
        expect(sortedStatuses).to.equal(statuses);
      });
  });
});
