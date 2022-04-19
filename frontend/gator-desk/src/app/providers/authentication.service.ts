import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {

  apiUrl = "http://localhost:8181/login"

  constructor(private http: HttpClient) { }

  continueLogin(userData : any){
    return this.http.post(this.apiUrl, userData);
  }
}
