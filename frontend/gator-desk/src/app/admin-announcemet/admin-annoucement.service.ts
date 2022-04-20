import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';


@Injectable()
export class AdminAnnouncementService {

  result: any;

  httpOptions = {
    withCredentials: true,
    headers: new HttpHeaders({ 
    "Content-Type": "application/x-www-form-urlencoded",
    "Accept": "*/*"
   })
  };

  constructor(private _http: HttpClient) { }

//   getMeetups() {
//     let headers = new Headers();
//     let options = new RequestOptions( {method: RequestMethod.Get, headers: headers });
//      return this._http.get("http://localhost:8000/getMeetups", options).map(result => this.result = result.json().data);
//   }

createAnnouncement(reqBody: any) {
     return this._http.post("http://localhost:8181/createAnnouncement",reqBody,this.httpOptions);
  }

}
