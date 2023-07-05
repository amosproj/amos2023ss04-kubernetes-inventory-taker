describe("Navigation", () => {
  it("has anchor tags using cy.get and .each", () => {
    //visit container list page
    cy.visit("/");

    //Find each a link and check href is defined
    cy.get("a").each(($a) => {
      const message = $a.parent().parent().text();
      expect($a, message).to.not.have.attr("href", "#undefined");
    });
  });
});
