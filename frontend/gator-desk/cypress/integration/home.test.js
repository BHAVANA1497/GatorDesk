
it('loads examples', () => {
    cy.visit('/');
    cy.contains(' What do you love? ')
    cy.contains(' Do more of it with Gator Meetup ')
  });


  it('tests navigation to annoucement component' , () => {
     cy.get('#annoucement').click();
     cy.contains('Announcements:')
  });

  it('tests navigation to admin annoucement component' , () => {
    cy.get('#login').click();
    cy.contains('Announcements:')
 });