const { type } = require("os");

it('loads examples', () => {
   cy.visit('/');
   cy.contains(' We hear you ')
   cy.contains(' Do more of it with Gator Desk ')
 });

//  it('should test login', () => {
//   cy.visit('/login');
//   cy.get('#email').type("Akhil2"); 
//   cy.get("#password").type("password"); 
//   cy.get('#loginid').click();

//   cy.request({
//     url: "http://localhost:8181/login",
//     method: 'POST',
//       body :{
//         username: 'Akhil2',
//         password: 'password',
//     }
//   }).then(res => {
//     localStorage.setItem('USER_DETAILS', res.body.username)
//     localStorage.setItem('isAdmin', res.body.IsAdmin)
//    });

//   cy.location("pathname").should("include", "/home");
//   cy.visit('/home')

//  });

 Cypress.Commands.add('login', (username, password) => {
  cy.visit('/login');

  cy.get('#email').type("Akhil2"); 

  cy.get("#password").type("password"); 

  cy.get('#loginid').click();
  cy.location("pathname").should("include", "/home");
})

 beforeEach(() => {
  cy.login('Akhil2', 'password')
})

 it('tests navigation to annoucement component' , () => {
  localStorage.getItem('USER_DETAILS')
  localStorage.getItem('isAdmin')
    cy.get('#annoucement').click();
    cy.intercept('GET', '/listAllAnnouncements')
    cy.contains('Announcements')
    cy.get("mat-form-field input").get(`[formcontrolname="eventTitle"]`).type("test test"); 
   cy.wait(2000)

   cy.get('#mat-select-0').click().get('mat-option').contains('Event').click();
   cy.wait(1000)
  cy.get("mat-form-field input").get(`[formcontrolname="eventDesc"]`).type("test test"); 
  cy.wait(2000)
  cy.get('.mat-button-wrapper').contains('Create Announcement').click()
  cy.intercept('POST', '/createAnnouncement')
  cy.wait(1000)
  cy.contains('Announcements')
  cy.visit('/home');
 }); 



it('tests navigation to Lost & Found component' , () => {
  localStorage.getItem('USER_DETAILS')
  localStorage.getItem('isAdmin')
    cy.get('#lostandfound').click({force: true});
    cy.wait(1000)
    cy.intercept('GET', '/listAllLostItems');
    cy.wait(2000)
    cy.intercept('GET', '/listAllFoundItems');
    cy.contains('Lost');
    cy.contains('Found');
 }); 


it('should test signup ', () => {
  cy.visit('/signup');
  cy.get('#name').type("TEST NAME"); 
  cy.get("#email").type("TEST EMAIL"); 
  cy.get('#password').type("TEST PASSWORD");
  cy.get("#signup").click();
  cy.visit('/home');

})