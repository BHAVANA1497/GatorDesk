it('loads examples', () => {
   cy.visit('/');
   cy.contains(' What do you love? ')
   cy.contains(' Do more of it with Gator Meetup ')
 });


 it('tests navigation to annoucement component' , () => {
    cy.get('#annoucement').click();
    cy.intercept('GET', '/listAllAnnouncements')
    cy.contains('Announcements:')
 });

 it('tests navigation to admin annoucement component' , () => {
   cy.get('#login').click();
   cy.contains('Announcements:')
});

it('Fill the annoucement details, post and check if it is getting posted', () => {
  cy.wait(2000)
  cy.get("mat-form-field input").get(`[formcontrolname="eventTitle"]`).type("test test"); 
  cy.wait(2000)

   cy.get('#mat-select-0').click().get('mat-option').contains('Event').click();
   cy.wait(1000)
  cy.get("mat-form-field input").get(`[formcontrolname="eventDesc"]`).type("test test"); 
  cy.wait(2000)
  cy.get('.mat-button-wrapper').contains('Create Announcement').click()
  cy.wait(1000)
  cy.contains('test test')
})