import { TestBed } from '@angular/core/testing';

import { AdminLostAndFoundService } from './admin-lost-and-found.service';

describe('AdminLostAndFoundService', () => {
  let service: AdminLostAndFoundService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AdminLostAndFoundService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
