import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { AdminAnnouncementService } from './admin-annoucement.service';

@Component({
  selector: 'app-admin-announcemet',
  templateUrl: './admin-announcement.component.html',
  styleUrls: ['./admin-announcement.component.css']
})
export class AdminAnnouncementComponent implements OnInit {
  announcementForm: FormGroup;
  eventTitle = '';
  eventCategory = '';
  eventDesc = '';
  postData: any = {};
  isAdmin: boolean = false;

  constructor(private fb: FormBuilder, public  admObj:AdminAnnouncementService ) { 
    this.announcementForm = fb.group({
      'eventTitle': [null, Validators.required],
      'eventDesc': [null, Validators.required]
    });
  }

  ngOnInit(): void {
    let isAdmin = localStorage.getItem("isAdmin");
    if (isAdmin) {
      console.log("is admin", isAdmin);
      this.isAdmin = true;
    }
  }

  selectedValue(selectedValue: any){
    this.eventCategory = selectedValue.value;
  }

  onSubmit(){

    this.eventTitle = this.announcementForm.value.eventTitle;
    this.eventDesc = this.announcementForm.value.eventDesc;
    console.log( this.eventTitle + " "+ this.eventCategory + " " + this.eventDesc);
    this.postData.announcementTitle = this.announcementForm.value.eventTitle;
    this.postData.eventDescription = this.announcementForm.value.eventDesc;
    this.postData.announcementCategory = this.eventCategory;

    this.admObj.createAnnouncement(this.postData).subscribe(res => {
      window.location.reload();
    });;
  
   
  }

}
