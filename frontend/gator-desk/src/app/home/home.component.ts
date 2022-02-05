import { Component, OnInit } from '@angular/core';
import { IvyCarouselModule } from 'angular-responsive-carousel';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  constructor(private router: Router) { 
    
  }

  public imageSources = [
    {path:'https://57b23a891d0ffc325f7b-559c437472c4da6a6044f5fd408e667f.ssl.cf1.rackcdn.com/Stoneridge-Apartments1475003638.jpg'},
    {path: 'https://cdngeneralcf.rentcafe.com/dmslivecafe/2/129655/50733CE0-0962-4BDD-A35D-81152816203B.jpeg?quality=85&scale=both&'},
    {path: 'https://ratemyapartments.s3.amazonaws.com/files/e7/a2/e7a2c601afbcb789cb2d464bf_original.jpg'},
    {path:'https://gatorrentals.com/wp-content/uploads/2017/03/original.0-25.jpeg'},
    {path:'https://57b23a891d0ffc325f7b-559c437472c4da6a6044f5fd408e667f.ssl.cf1.rackcdn.com/StoneridgeApartments1343849872.jpg'}
  ];

  public config: IvyCarouselModule = {
    animation: true,
    autoplay: true,
    autoplayDelay: 2000,
    stopAutoplayMinWidth: 768
  };

  navigate(){
    this.router.navigate(['annoucement']);
    }
 

  ngOnInit(): void {
  }

}
