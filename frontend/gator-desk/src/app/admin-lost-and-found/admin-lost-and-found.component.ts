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
  isAdmin: boolean = false;

  constructor(public adminService: AdminLostAndFoundService) {
    this.getLostItems();
    this.getFoundItems();
  }

  ngOnInit(): void {
    let isAdmin = localStorage.getItem("isAdmin");
    if (isAdmin) {
      this.isAdmin = true;
    }
  }

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
    let lostObj = this.getFoundItem(foundId);
    lostObj.found_id =  foundId;
    this.adminService.linkFoundItem( lostId, lostObj).subscribe((data)=> {
      console.log(data);
    });
    this.getLostItems();
    this.getFoundItems();
  }

  getLostItem( lostItemId : any) : any {
    let lostRec : any;
    this.adminService.getLostItem(lostItemId).subscribe((data) => {
      lostRec = data.data;
    })
    return lostRec;
  }

  getFoundItem(foundId: any) : any{
    let resp: any;
    this.adminService.getFoundItem(foundId).subscribe((data) => {
      console.log(data);
      resp = data;
    });

    console.log(resp);
    return resp;
  }
}
