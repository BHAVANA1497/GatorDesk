import { Component, OnInit } from '@angular/core';
import { AdminLostAndFoundService } from './admin-lost-and-found.service';

@Component({
  selector: 'app-admin-lost-and-found',
  templateUrl: './admin-lost-and-found.component.html',
  styleUrls: ['./admin-lost-and-found.component.css'],
})
export class AdminLostAndFoundComponent implements OnInit {
  lostItems$: any;
  foundItems$: any;
  currentFoundItem$: any;

  constructor(public adminService: AdminLostAndFoundService) {
    this.getLostItems();
    this.getFoundItems();
  }

  ngOnInit(): void {}

  getLostItems() {
    this.adminService.getLostItems().subscribe((data) => {
      console.log(data);
      this.lostItems$ = data;
    });
  }

  getFoundItems() {
    this.adminService.getFoundItems().subscribe((data) => {
      console.log(data);
      this.foundItems$ = data;
    });
    this.lostItems$.forEach((item: any) => {
      if (item.isFound) {
        item.foundItem = this.getFoundItem(item.foundId);
      }
    });
    console.log(this.lostItems$);
  }

  linkFoundItem(foundId: any, lostId: any) {
    console.log(foundId + ' = ' + lostId);
    this.adminService.linkFoundItem(foundId, lostId);

    this.getLostItems();
    this.getFoundItems();
  }

  getFoundItem(foundId: any) {
    let resp: any;
    this.adminService.getFoundItem(foundId).subscribe((data) => {
      console.log(data);
      resp = data;
    });

    console.log(resp);
    return resp;
  }
}
