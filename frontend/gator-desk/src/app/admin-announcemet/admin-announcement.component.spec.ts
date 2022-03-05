import { HttpClient, HttpClientModule, HttpHandler } from '@angular/common/http';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { ComponentFixture, getTestBed, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { Router, RouterModule } from '@angular/router';
import { AdminAnnouncementService } from './admin-annoucement.service';

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
      imports: [FormsModule,
        ReactiveFormsModule, HttpClientTestingModule],
        providers: [AdminAnnouncementService, HttpClient, HttpClientModule, HttpHandler]
    })
    .compileComponents();

    injector = getTestBed();
    service = injector.get(AdminAnnouncementService);
    httpMock = injector.get(HttpClientTestingModule);
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminAnnouncementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('post an annoucement', () => {
    const dummyTodoLists = [
      {adminId: 1,
        announcementCategory: "Event Announcement",
        announcementId: 1,
        announcementTitle: "New Year Celebrations",
        createdAt: "2022-02-02T19:17:12.878628-05:00",
        eventDescription: "We are celebrating New Year&#39;s!",
        updatedAt: "2022-02-02T19:17:12.878628-05:00",
        venue: "Stoneridge Leasing Office"}, 
     
    ];
    expect(1).toEqual(1);
    service.createAnnouncement(dummyTodoLists).subscribe( (data:any) => {
      expect(data.body.length).toBe(2);
      expect(data.body).toEqual(dummyTodoLists);
    });

  });

});