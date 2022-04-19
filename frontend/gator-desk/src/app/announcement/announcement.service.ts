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
    "credentials": "include",
    "Accept": "/"
   })
  };

  constructor(private _http: HttpClient) { }

 getAllAnnouncements(): Observable<any> {
      return this._http.get("http://localhost:8181/listAllAnnouncements",this.httpOptions);
  }

}