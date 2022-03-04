import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';


@Injectable()
export class AnnouncemetService {

  annoucements: Object | undefined;
  result: any;

  httpOptions = {
    headers: new HttpHeaders({ 
    "Content-Type": "application/json",
    "Accept": "*/*"
   })
  };

  constructor(private _http: HttpClient) { }

//   getMeetups() {
//     let headers = new Headers();
//     let options = new RequestOptions( {method: RequestMethod.Get, headers: headers });
//      return this._http.get("http://localhost:8000/getMeetups", options).map(result => this.result = result.json().data);
//   }

 getAllAnnouncements(): Observable<any> {
      return this._http.get("http://localhost:8181/listAllAnnouncements",this.httpOptions);
  }

}
