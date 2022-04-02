import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LostAndFoundService } from './lost-and-found.service';

@Component({
  selector: 'app-lost-and-found',
  templateUrl: './lost-and-found.component.html',
  styleUrls: ['./lost-and-found.component.css']
})
export class LostAndFoundComponent implements OnInit {

  lostAndFoundForm: FormGroup;
  losttype = '';
  eventCategory = '';
  desc = '';
  details = '';
  postData: any = {};
  lostItems$: any;
  foundItems$: any;

  constructor(private fb: FormBuilder , public lostservice: LostAndFoundService) { 
    this.lostAndFoundForm = fb.group({
      'losttype': [],
      'desc': [null, Validators.required],
      'details': []
    });
    this.getLostItems();
    this.getFoundItems();
  }

  getLostItems() {
    this.lostservice.getLostItems().subscribe((data) => {
      console.log(data);
      this.lostItems$ = data;
    });
  }

  getFoundItems() {
    this.lostservice.getFoundItems().subscribe((data) => {
      console.log(data);
      this.foundItems$ = data;
    });
    console.log(this.lostItems$);
  }

  ngOnInit(): void {
  }

  selectedValue(selectedValue: any){
    this.eventCategory = selectedValue.value;
  }

  onSubmit(){
    this.losttype = this.lostAndFoundForm.value.losttype;
    this.desc = this.lostAndFoundForm.value.eventDesc;
    console.log( this.losttype + " "+ this.eventCategory + " " + this.desc);
    // this.postData.announcementTitle = this.lostAndFoundForm.value.losttype;
    // this.postData.eventDescription = this.lostAndFoundForm.value.desc;
    // this.postData.announcementCategory = this.eventCategory;

    // this.admObj.createAnnouncement(this.postData).subscribe(res => {
    //   
    // });
    window.location.reload();
  }

}
