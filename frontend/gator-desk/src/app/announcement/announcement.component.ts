import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { AnnouncemetService } from './announcement.service';


@Component({
  selector: 'app-announcement',
  templateUrl: './announcement.component.html',
  styleUrls: ['./announcement.component.css']
})
export class AnnouncementComponent implements OnInit {
  panelOpenState = false;
  annoucements$ : any;
  constructor(public announcementService: AnnouncemetService) { 
   this.announcementService.getAllAnnouncements().subscribe((data)=>{
     console.log(data.data);
   this.annoucements$ = data.data;
   })
  
  }

  ngOnInit(): void {
  }

}