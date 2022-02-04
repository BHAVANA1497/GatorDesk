import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminAnnouncemetComponent } from './admin-announcemet.component';

describe('AdminAnnouncemetComponent', () => {
  let component: AdminAnnouncemetComponent;
  let fixture: ComponentFixture<AdminAnnouncemetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AdminAnnouncemetComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminAnnouncemetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
