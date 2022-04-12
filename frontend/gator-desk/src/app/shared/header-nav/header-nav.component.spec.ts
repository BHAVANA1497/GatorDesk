import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { UserService } from 'src/app/providers/user.service';

import { HeaderNavComponent } from './header-nav.component';

describe('HeaderNavComponent', () => {
  let component: HeaderNavComponent;
  let fixture: ComponentFixture<HeaderNavComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HeaderNavComponent ],
      providers: [UserService],
      imports: [HttpClientModule, RouterTestingModule, HttpClientTestingModule]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HeaderNavComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should render title in a h1 tag', () => {
    const fixture = TestBed.createComponent(HeaderNavComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h1').textContent).toContain('GatorDesk');
  });

  
  it('should have Login button enabled', () => {
    const btn = fixture.debugElement.nativeElement.querySelector('#login');
    expect(btn.innerHTML).toBeTruthy();
  });

  it('should have Signup button enabled', () => {
    const btn = fixture.debugElement.nativeElement.querySelector('#signup');
    expect(btn.innerHTML).toBeTruthy();
  });

  it('should have maintenancerequest button enabled', () => {
    const btn = fixture.debugElement.nativeElement.querySelector('#maintenancerequest');
    expect(btn.innerHTML).toBeTruthy();
  });

  it('should have lostandfound button enabled', () => {
    const btn = fixture.debugElement.nativeElement.querySelector('#lostandfound');
    expect(btn.innerHTML).toBeTruthy();
  });

  it('should have annoucement button enabled', () => {
    const btn = fixture.debugElement.nativeElement.querySelector('#annoucement');
    expect(btn.innerHTML).toBeTruthy();
  });

  it('should have home button enabled', () => {
    const btn = fixture.debugElement.nativeElement.querySelector('#home');
    expect(btn.innerHTML).toBeTruthy();
  });



});
