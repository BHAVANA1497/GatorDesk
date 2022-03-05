import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';
import { AdminAnnouncementService } from './admin-annoucement.service';
import {
  HttpClientTestingModule,
  HttpTestingController
} from '@angular/common/http/testing';


import { AdminAnnouncementComponent } from './admin-announcement.component';

describe('AdminAnnouncementComponent', () => {
  let component: AdminAnnouncementComponent;
  let fixture: ComponentFixture<AdminAnnouncementComponent>;
  let injector: TestBed;
  let service: AdminAnnouncementService;
  let httpMock: HttpTestingController;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AdminAnnouncementComponent ],
      imports: [HttpClientModule, FormsModule, ReactiveFormsModule, BrowserModule , RouterModule, RouterTestingModule],
      providers: [AdminAnnouncementService]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminAnnouncementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

});
