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
  isAdmin = 'user';

  displayedColumnsLost = ['lost_type','CreatedAt' ,'description'];
  displayedColumnsFound = ['lost_type','CreatedAt', 'description'];

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
    // this.getLostItems();
    // this.getFoundItems();
  }

  getLostItems() {
    this.lostservice.getLostItems().subscribe((data) => {
      console.log(data.data);
      this.lostItems$ = data.data;
      this.datasourceLost = new MatTableDataSource(this.lostItems$);
      this.datasourceLost.paginator = this.paginatorLost;
    });
  }

  getFoundItems() {
    this.lostservice.getFoundItems().subscribe((data) => {
      console.log(data.data);
      this.foundItems$ = data.data;
      this.datasourceFound = new MatTableDataSource(this.foundItems$);
      this.datasourceFound.paginator = this.paginatorFound;
    });
  }

  ngOnInit(): void {
    let isAdmin = localStorage.getItem("isAdmin");
    if (isAdmin === 'true') {
      this.isAdmin = 'admin';
    }
    this.lostservice.getLostItems().subscribe((data) => {
      console.log(data.data);
      this.lostItems$ = data.data;
      this.datasourceLost = new MatTableDataSource(this.lostItems$);
      this.datasourceLost.paginator = this.paginatorLost;
      console.log(this.datasourceLost);
    });
    this.lostservice.getFoundItems().subscribe((data) => {
      console.log(data.data);
      this.foundItems$ = data.data;
      this.datasourceFound = new MatTableDataSource(this.foundItems$);
      this.datasourceFound.paginator = this.paginatorFound;
      console.log(this.datasourceFound);
    });
  }

  selectedRequestType(selectedValue: any){
    this.requestType = selectedValue.value;
  }

  selectedLostType( selectedValue: any){
     this.lostType = selectedValue.value;
  }

  onSubmit(){
    console.log(this.lostType);
    this.lostType = this.lostType;
    this.requestType = this.requestType;
    this.desc = this.lostAndFoundForm.value.desc;

    this.postDataLost.lost_type = this.lostType;
    this.postDataLost.description = this.lostAndFoundForm.value.desc;
 
    this.postDataFound.found_type = this.lostType;
    this.postDataFound.description = this.lostAndFoundForm.value.desc;

  
    if( this.requestType == "lost"){
      this.lostservice.createLostRequest(this.postDataLost).subscribe(res => {
        console.log(res);
       window.location.reload();
        
      });
    }
    if( this.requestType == "found"){
      this.lostservice.createFoundRequest(this.postDataFound).subscribe(res => {
        console.log(res);
       window.location.reload();
        
      });
    }
  }

}
