import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SignupService {

  
  httpOptions = {
    headers: new HttpHeaders({ 
    "Content-Type": "application/json",
    "Accept": "*/*"
   })
  };

  constructor(private _http: HttpClient) { }

  createUser(reqBody: any) {
    return this._http.post("http://localhost:8181/signUp",reqBody,this.httpOptions);
 }

}
