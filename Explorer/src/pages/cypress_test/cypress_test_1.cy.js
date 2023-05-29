import CypressTest1 from "./cypress_test_1.tsx";

describe("<CypressTest1 />", () => {
  it("should render and display expected content", () => {
    // Mount the React component for the Cypress Test page
    cy.mount(<CypressTest1 />);

    // The new page should contain an h1 with "Cypress Test 1"
    cy.get("h1").contains("Cypress Test 1");

    // Validate that a link with the expected URL is present
    // *Following* the link is better suited to an E2E test
    cy.get('a[href="/cypress_test/cypress_test_2"]').should("be.visible");
  });
});
