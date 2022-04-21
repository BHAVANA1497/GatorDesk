import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {

  apiUrl = "http://localhost:8181/login";

  httpOptions = {
    headers: new HttpHeaders({ 
    "Content-Type": "application/x-www-form-urlencoded",
    "Accept": "*/*"
   })
  };


  constructor(private http: HttpClient) { }

  continueLogin(userData : any){
    return this.http.post(this.apiUrl, userData, this.httpOptions);
  }

  isLoggedIn(){
    return localStorage.getItem('USER_DETAILS') != null;
  }
}
