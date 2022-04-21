import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-maintenance-request',
  templateUrl: './maintenance-request.component.html',
  styleUrls: ['./maintenance-request.component.css']
})
export class MaintenanceRequestComponent implements OnInit {

  isAdmin: string = 'user';

  constructor() { }

  ngOnInit(): void {
    let isAdmin = localStorage.getItem("isAdmin");
   if (isAdmin === 'true') {
      this.isAdmin = 'admin';
    }
  }

}
