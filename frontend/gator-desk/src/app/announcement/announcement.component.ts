import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
<<<<<<< HEAD
import { AnnouncemetService } from './annoucement.service';
=======
import { AnnouncemetService } from './announcement.service';
>>>>>>> 8d67ae27258b2ef845674be04d11f290d617135a


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