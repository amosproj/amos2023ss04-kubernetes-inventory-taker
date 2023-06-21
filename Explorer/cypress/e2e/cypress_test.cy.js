describe("Navigation", () => {
  it("should navigate to the about page", () => {
    // Start from cypress_test_1

    cy.visit("/cypress_test/cypress_test_1");

    // Find a link with an href attribute containing "cypress_test_2" and click it
    cy.get('a[href*="cypress_test_2"]').click();

    // The new url should include "/cypress_test_2"
    cy.url().should("include", "/cypress_test_2");

    // The new page should contain an h1 with "Cypress Test 1"
    cy.get("h1").contains("Cypress Test 2");
  });
});
