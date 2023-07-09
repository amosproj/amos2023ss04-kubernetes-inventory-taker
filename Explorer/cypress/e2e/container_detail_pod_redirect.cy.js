describe("Navigation", () => {
  it("navigates to a random container detail page and checks pod_id link", () => {
    // Visit container list page
    cy.visit("/containers");

    // Collect all container links
    cy.get('a[href*="/containers/"]').then(($links) => {
      // Choose a random index
      const randomIndex = Math.floor(Math.random() * $links.length);
      // Get the link at the random index
      const randomLink = $links[randomIndex];
      // Click the random link to navigate to container detail page
      cy.wrap(randomLink).click();
    });

    // Once on the container detail page, find the link in the row named 'POD_ID' and click on it
    cy.contains("td", "POD_ID").parent().find("a").click();

    // Assert that the URL is correct
    cy.url().should("include", "/pods/");
  });
});
