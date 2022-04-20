import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthenticationService } from '../providers/authentication.service';
import { UserService } from '../providers/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  userData = {
    username: '',
    password: '',
  };

  responseData: any;

  constructor(private _http: HttpClient, public authservice: AuthenticationService,  private router: Router) { }

  ngOnInit(): void {
    
  }

  doLogin() {
    this.authservice.continueLogin(this.userData).subscribe( res =>{
      console.log(res);  
      this.responseData = res;
      localStorage.setItem('USER_DETAILS', this.userData.username);
      localStorage.setItem('isAdmin', this.responseData.IsAdmin)
       this.router.navigate(['/home']);
    }, err => {
      console.log(err.error);
    });
    
  }

}
