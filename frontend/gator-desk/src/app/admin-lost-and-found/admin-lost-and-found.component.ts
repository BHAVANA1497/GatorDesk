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
  isAdmin = 'user';
  lostRec : any;

  constructor(public adminService: AdminLostAndFoundService) {
    this.getLostItems();
    this.getFoundItems();
  }

  ngOnInit(): void {
    let isAdmin = localStorage.getItem("isAdmin");
    if (isAdmin) {
      this.isAdmin = 'admin';
    }
  }

  getLostItems() {
    this.adminService.getLostItems().subscribe((data) => {
      console.log(data);
      this.lostItems$ = data.data;
    });
  }

  getFoundItems() {
    this.adminService.getFoundItems().subscribe((data) => {
      console.log(data);
      this.foundItems$ = data.data;
    });
  }

  linkFoundItem(foundId: any, lostId: any) {
    console.log(foundId + ' = ' + lostId);
    let lostRec : any;
    let fId : number = Number(foundId);
    this.adminService.getLostItem(lostId).subscribe((data) => {
      console.log(data);
      lostRec = data.data[0];
      this.lostRec = lostRec;
      this.lostRec.found_id =  fId;
      this.lostRec.is_found = true;
      console.log(lostRec);

    this.adminService.linkFoundItem( lostId, lostRec).subscribe((data)=> {
      console.log(data);
      window.location.reload();
    });

    });
  }

  getLostItem( lostItemId : any) : any {
    let lostRec : any;
    this.adminService.getLostItem(lostItemId).subscribe((data) => {
      console.log(data);
      lostRec = data;
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
