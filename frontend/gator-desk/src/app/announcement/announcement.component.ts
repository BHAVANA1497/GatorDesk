import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AnnouncemetService } from './annoucement.service';


@Component({
  selector: 'app-announcement',
  templateUrl: './announcement.component.html',
  styleUrls: ['./announcement.component.css']
})
export class AnnouncementComponent implements OnInit {
  panelOpenState = false;
  annoucements$ : any;
  isAdmin = 'user';

  constructor(public announcementService: AnnouncemetService, private router: Router) { 
   this.announcementService.getAllAnnouncements().subscribe((data)=>{
     console.log(data.data);
   this.annoucements$ = data.data;
   })
  
  }

  ngOnInit(): void {
    let isAdmin = localStorage.getItem("isAdmin");
    if (isAdmin === 'true') {
      this.isAdmin = 'admin';
    }
  }

}