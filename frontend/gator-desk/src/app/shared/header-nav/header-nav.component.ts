import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from 'src/app/providers/user.service';

@Component({
  selector: 'app-header-nav',
  templateUrl: './header-nav.component.html',
  styleUrls: ['./header-nav.component.css']
})
export class HeaderNavComponent implements OnInit {

  isLoggedIn = false;

  constructor(
    public user: UserService,
    private _http: HttpClient,
    private router: Router) {
    if (this.user._token) {
      this.loginCheck();
    }
  }

  ngOnInit(): void {
    let user = localStorage.getItem("USER_DETAILS");
    console.log("%o ", user);
    //this.name = user.name;
    if (user != null) {
      this.isLoggedIn = true;
    }
  }

  loginCheck() {
    console.log('tes');
    this._http.post('users/islogin', { token: this.user._token })
      .subscribe(res => {
        console.log(res);
        //this.user._details = res.data;
      }, err => {
        console.log(err);
        this.logout();
      });
  }

  logout() {
    this.user._details = null;
    this.user._token = '';
    localStorage.clear();
    this.router.navigate(['/login']);
  }

}
