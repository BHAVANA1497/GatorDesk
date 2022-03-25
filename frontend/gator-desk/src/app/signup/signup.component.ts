import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { UserService } from '../providers/user.service';
import { SignupService } from './signup.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {
  userData = {
    firstname: '',
    username: '',
    password: ''
  };
  constructor( private _http: HttpClient, public signup: SignupService,  private router: Router) { 
  }

  ngOnInit(): void {
  }

  doSignup(){
    console.log(this.userData);
    this.signup.createUser(this.userData).subscribe(res => {
      console.log(res);
      alert('User has been created, you can login with your credentials');
      this.router.navigate(['/home']);
    }, err => {
      console.log(err.error);
      alert(err.error);
    });
  }

}
