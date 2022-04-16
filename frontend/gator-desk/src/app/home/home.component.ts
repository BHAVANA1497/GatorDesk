import { Component, OnInit } from '@angular/core';
import { IvyCarouselModule } from 'angular-responsive-carousel';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
})
export class HomeComponent implements OnInit {
  public imageSources = [
    {path: '../../assets/slider2.jpg'},
    {path: '../../assets/desk.jpg'},
    {path: '../../assets/apartments.jpg'},
    
  ];
  
  isAdmin = false;

  constructor() {}

  public config: IvyCarouselModule = {
    animation: true,
    autoplay: true,
    autoplayDelay: 2000,
    stopAutoplayMinWidth: 768,
  };

  ngOnInit(): void {
    let isAdmin = localStorage.getItem("isAdmin");
    console.log("is admin", isAdmin);
    if (isAdmin) {  
      this.isAdmin = true;
    }
  }

}
