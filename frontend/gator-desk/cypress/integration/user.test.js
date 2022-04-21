const { type } = require("os");

it('loads examples', () => {
   cy.visit('/');
   cy.contains(' We hear you ')
   cy.contains(' Do more of it with Gator Desk ')
 });

 Cypress.Commands.add('login', (username, password) => {
  cy.visit('/login');
  cy.get('#email').type("user2"); 
  cy.get("#password").type("user2"); 
  cy.get('#loginid').click();
  cy.location("pathname").should("include", "/login");
})

 beforeEach(() => {
  cy.login('user2', 'user2')
})

 it('tests navigation to annoucement component' , () => {
  cy.wait(1000)
    cy.get('#annoucement').click();
    cy.wait(1000)
    cy.intercept('GET', '/listAllAnnouncements')
    cy.wait(1000)
    cy.contains('Announcements')
 }); 

 it('Fill Lost and Found details, post and check if it is getting posted', () => {
  cy.contains('Lost');
  cy.contains('Found');
  cy.get('#lostandfound').click({force: true});
  cy.wait(1000)
  cy.get('mat-select[formControlName=requestType]').click().get('mat-option').contains('LOST').click();
  cy.wait(2000)
  cy.get('mat-select[formControlName=itemType]').click().get('mat-option').contains('Book').click(); 
  cy.wait(2000)
  cy.get(`[formcontrolname="desc"]`).type("test test"); 
  cy.wait(2000)
  cy.get(`[formcontrolname="details"]`).type("test test"); 
  cy.wait(2000)
  cy.get('.mat-button-wrapper').contains('Create Request').click({force: true})
  cy.wait(1000)
  cy.contains('Found')
  cy.visit('/home');
})
