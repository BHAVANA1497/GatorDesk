import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../providers/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  userData = {
    email: '',
    password: '',
  };
  constructor(private _http: HttpClient, public user: UserService,  private router: Router) { }

  ngOnInit(): void {
  }

  doLogin() {
    console.log(this.userData);
    if(this.userData.email == "test@gmail.com" && this.userData.password == "test"){
      this.user._details = this.userData;
      localStorage.setItem('USER_DETAILS', JSON.stringify(this.userData));
      this.user._token = "dummytoken";
    }
    this.router.navigate(['/home']);
    // this._http.post('users/login', this.userData)
    //   .subscribe(res => {
    //     console.log(res);
    //     // this.user._details = res['data'];
    //     // this.user._token = res['token'];
    //     // localStorage.setItem('TOKEN', res['token']);
    //     // localStorage.setItem('USER_DETAILS', JSON.stringify(res['data']));
    //     this.router.navigate(['/home']);
    //     // alert(res.message);
    //   }, err => {
    //     console.log(err);
    //     alert(JSON.parse(err._body).message);
    //   });
  }

}
