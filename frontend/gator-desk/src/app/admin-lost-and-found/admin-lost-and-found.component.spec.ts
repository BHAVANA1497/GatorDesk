import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { AdminLostAndFoundComponent } from './admin-lost-and-found.component';
import { AdminLostAndFoundService } from './admin-lost-and-found.service';

describe('AdminLostAndFoundComponent', () => {
  let component: AdminLostAndFoundComponent;
  let fixture: ComponentFixture<AdminLostAndFoundComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AdminLostAndFoundComponent ],
      providers: [AdminLostAndFoundService],
      imports: [HttpClientModule, RouterTestingModule, HttpClientTestingModule]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminLostAndFoundComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
