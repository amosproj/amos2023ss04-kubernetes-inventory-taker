describe("Template", () => {
  beforeEach(() => {
    cy.visit("http://localhost:3000/");
  });

  it("url should be/include/contain ...", () => {
    cy.url().should("eq", "http://localhost:3000/");
    cy.url().should("include", "3000");
    cy.url().should("contain", "3000");
  });

  // data-cy
  it("should have a table > data-cy", () => {
    cy.get('[data-cy="container-table"]').should("be.visible");
  });
  // table
  it("should have a table > table", () => {
    cy.get("table").should("be.visible");
  });
  // css
  it("should have a table > css", () => {
    cy.get(".w-full.text-left.text-sm").should("be.visible");
  });

  it("table contains 11 rows", () => {
    cy.get('[data-cy="container-table"] > tbody')
      .children()
      .should("have.length", 11);
  });

  it("table should contain 'nginx'", () => {
    cy.get('[data-cy="container-table"] > tbody > tr > td > a').contains(
      "nginx"
    );
  });

  it("second row has 'etcd'", () => {
    // second tr inside (table) tbody
    cy.get("tbody > tr").eq(1).find("td > a").should("have.text", "etcd");
  });

  it("clicking on name opens toast", () => {
    expect(Cypress.$('button[aria-label="Close"]')).not.to.exist;
    cy.contains("td", "etcd").find("a").click();
    cy.get('button[aria-label="Close"]').parent().should("be.visible");
  });
});
