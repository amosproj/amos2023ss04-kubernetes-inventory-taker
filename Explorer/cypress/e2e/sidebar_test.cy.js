describe("Sidebar", () => {
  it("Should navigate to the dashboard page", () => {
    cy.visit("http://localhost:3000/");
    cy.get('a[href*="dashboard"]').click();
    cy.url().should("include", "/dashboard");
    cy.get("h1").contains("Dashboard");
  });

  it("Should navigate to the cluster page", () => {
    cy.visit("http://localhost:3000/");
    cy.get('a[href*="cluster"]').click();
    cy.url().should("include", "/cluster");
    cy.get("h1").contains("Cluster");
  });

  it("Should navigate to the deployments page", () => {
    cy.visit("http://localhost:3000/");
    cy.get('a[href*="deployments"]').click();
    cy.url().should("include", "/deployments");
    cy.get("h1").contains("Deployments");
  });

  it("Should navigate to the nodes page", () => {
    cy.visit("http://localhost:3000/");
    cy.get('a[href*="nodes"]').click();
    cy.url().should("include", "/nodes");
    cy.get("h1").contains("Nodes");
  });

  it("Should navigate to the pods page", () => {
    cy.visit("http://localhost:3000/");
    cy.get('a[href*="pods"]').click();
    cy.url().should("include", "/pods");
    cy.get("h1").contains("Pods");
  });

  it("Should navigate to the containers page", () => {
    cy.visit("http://localhost:3000/");
    cy.get('a[href*="containers"]').click();
    cy.url().should("include", "/containers");
    cy.get("h1").contains("Containers");
  });

  it("Should navigate to the services page", () => {
    cy.visit("http://localhost:3000/");
    cy.get('a[href*="services"]').click();
    cy.url().should("include", "/services");
    cy.get("h1").contains("Services");
  });

  it("Should navigate to the volumes page", () => {
    cy.visit("http://localhost:3000/");
    cy.get('a[href*="volumes"]').click();
    cy.url().should("include", "/volumes");
    cy.get("h1").contains("Volumes");
  });

  it("Should navigate to the settings page", () => {
    cy.visit("http://localhost:3000/");
    cy.get('a[href*="settings"]').click();
    cy.url().should("include", "/settings");
    cy.get("h1").contains("Settings");
  });
});
