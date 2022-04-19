import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-admin-header-nav',
  templateUrl: './admin-header-nav.component.html',
  styleUrls: ['./admin-header-nav.component.css']
})
export class AdminHeaderNavComponent implements OnInit {

  constructor(private router: Router) { }
  isLoggedIn = false;
  userName  = '';


  ngOnInit(): void {
    let userName = localStorage.getItem("USER_DETAILS");
    if (userName != null) {
      this.isLoggedIn = true;
      this.userName = userName;
    }
  }

  
  logout() {
    localStorage.clear();
    this.router.navigate(['/login']);
  }

}
