import { HttpClient, HttpClientModule, HttpHandler } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { Router, RouterModule } from '@angular/router';
import { AdminAnnouncementService } from './admin-annoucement.service';

import { AdminAnnouncementComponent } from './admin-announcement.component';

describe('AdminAnnouncementComponent', () => {
  let component: AdminAnnouncementComponent;
  let fixture: ComponentFixture<AdminAnnouncementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AdminAnnouncementComponent ],
      imports: [FormsModule,
        ReactiveFormsModule, RouterModule, Router],
        providers: [AdminAnnouncementService, HttpClient, HttpClientModule, HttpHandler]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminAnnouncementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

});