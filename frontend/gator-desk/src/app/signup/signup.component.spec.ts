import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { FormsModule } from '@angular/forms';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { UserService } from '../providers/user.service';

import { SignupComponent } from './signup.component';
import { By } from '@angular/platform-browser';

describe('SignupComponent', () => {
  let component: SignupComponent;
  let fixture: ComponentFixture<SignupComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SignupComponent ],
      providers: [UserService],
      imports: [HttpClientModule,RouterTestingModule,HttpClientTestingModule, FormsModule]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SignupComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render data in a h4 tag', () => {
    const fixture = TestBed.createComponent(SignupComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h4').textContent).toContain('Signup');
  });

  it('should have Signup button enabled', () => {
    const btn = fixture.debugElement.nativeElement.querySelector('#signup');
    expect(btn.innerHTML).toBeTruthy();
  });

  it('should have password', () => {
    const board = fixture.debugElement.query(By.css('.top25')).nativeElement;
    expect(board.innerHTML).not.toBeNull();
    expect(board.innerHTML.length).toBeGreaterThan(0);
  });

  it('should have name', () => {
    const board = fixture.debugElement.query(By.css('.top25')).nativeElement;
    expect(board.innerHTML).not.toBeNull();
    expect(board.innerHTML.length).toBeGreaterThan(0);
  });

  it('should have email', () => {
    const board = fixture.debugElement.query(By.css('.top25')).nativeElement;
    expect(board.innerHTML).not.toBeNull();
    expect(board.innerHTML.length).toBeGreaterThan(0);
  });

  it('should render content from a tag', () => {
    const fixture = TestBed.createComponent(SignupComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('a').textContent).toContain('Already have an account? Login');
  });

});
