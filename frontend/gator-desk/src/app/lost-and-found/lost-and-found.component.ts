import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { LostAndFoundService } from './lost-and-found.service';

@Component({
  selector: 'app-lost-and-found',
  templateUrl: './lost-and-found.component.html',
  styleUrls: ['./lost-and-found.component.css']
})
export class LostAndFoundComponent implements OnInit {

  lostAndFoundForm: FormGroup;

  eventCategory = '';
  desc = '';
  requestType = '';
  lostType = '';
  postDataLost: any = {};
  postDataFound: any = {};
  lostItems$: any;
  foundItems$: any;
  isAdmin: boolean = false;

  displayedColumnsLost = ['lost_type', 'description'];
  displayedColumnsFound = ['lost_type', 'description'];

  datasourceLost !: MatTableDataSource<any>;
  @ViewChild('paginator') paginatorLost!: MatPaginator;

  datasourceFound !: MatTableDataSource<any>;
  @ViewChild('paginator') paginatorFound!: MatPaginator;

  constructor(private fb: FormBuilder , public lostservice: LostAndFoundService) { 
    this.lostAndFoundForm = fb.group({
      'requestType': [null, Validators.required],
      'desc': [null, Validators.required],
      'itemType':  [null, Validators.required],
      'details': []
    });
    this.getLostItems();
    this.getFoundItems();
  }

  getLostItems() {
    this.lostservice.getLostItems().subscribe((data) => {
      console.log(data.data);
      this.lostItems$ = data.data;
    });
  }

  getFoundItems() {
    this.lostservice.getFoundItems().subscribe((data) => {
      console.log(data.data);
      this.foundItems$ = data.data;
    });
    console.log(this.lostItems$);
  }

  ngOnInit(): void {
    let isAdmin = localStorage.getItem("isAdmin");
    if (isAdmin) {
      this.isAdmin = true;
    }
    this.lostservice.getLostItems().subscribe((data) => {
      console.log(data.data);
      this.lostItems$ = data.data;
      this.datasourceLost = new MatTableDataSource(this.lostItems$);
      this.datasourceLost.paginator = this.paginatorLost;
    });
    this.lostservice.getFoundItems().subscribe((data) => {
      console.log(data.data);
      this.foundItems$ = data.data;
      this.datasourceFound = new MatTableDataSource(this.foundItems$);
      this.datasourceFound.paginator = this.paginatorFound;
    });
  }

  selectedRequestType(selectedValue: any){
    this.requestType = selectedValue.value;
  }

  selectedLostType( selectedValue: any){
     this.lostType = selectedValue.value;
  }

  onSubmit(){
    this.lostType = this.lostType;
    this.requestType = this.requestType;
    this.desc = this.lostAndFoundForm.value.desc;
    this.postDataLost.lostType = this.lostType;
    this.postDataLost.desc = this.lostAndFoundForm.value.desc;
 
    this.postDataFound.FoundType = this.lostType;
    this.postDataFound.desc = this.lostAndFoundForm.value.desc;

   console.log(this.requestType);
  
    if( this.requestType == "lost"){
      this.lostservice.createLostRequest(this.postDataLost).subscribe(res => {
        console.log(res);
        
      });
    }
    if( this.requestType == "found"){
      this.lostservice.createFoundRequest(this.postDataFound).subscribe(res => {
        console.log(res);
        
      });
    }
  }

}
