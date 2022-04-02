import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../providers/user.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {
  userData = {
    name: '',
    email: '',
    password: ''
  };

  constructor(private _http: HttpClient, public user: UserService,  private router: Router) { }

  ngOnInit(): void {
  }

  doSignup(){
    console.log(this.userData);
    this.router.navigate(['/home']);
  }

}
