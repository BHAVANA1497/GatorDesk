import { HttpClientModule, HttpEvent, HttpEventType} from '@angular/common/http';
import { ComponentFixture, getTestBed, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AnnouncemetService } from './annoucement.service';
import { AnnouncementComponent } from './announcement.component';
import {
  HttpClientTestingModule,
  HttpTestingController
} from '@angular/common/http/testing';


describe('AnnouncementComponent', () => {
  let component: AnnouncementComponent;
  let fixture: ComponentFixture<AnnouncementComponent>;
  let injector: TestBed;
  let service: AnnouncemetService;
  let httpMock: HttpTestingController;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AnnouncementComponent ],
      providers: [AnnouncemetService],
      imports: [HttpClientModule, RouterTestingModule, HttpClientTestingModule]
    })
    .compileComponents();

    injector = getTestBed();
    service = injector.get(AnnouncemetService);
    httpMock = injector.get(HttpClientTestingModule);
  });




  it('should return list of Annoucemements>', () => {
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
    service.getAllAnnouncements().subscribe( (data:any) => {
      expect(data.body.length).toBe(2);
      expect(data.body).toEqual(dummyTodoLists);
    });

  });


});