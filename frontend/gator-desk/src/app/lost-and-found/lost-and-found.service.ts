import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { of } from 'rxjs/internal/observable/of';

@Injectable({
  providedIn: 'root',
})
export class LostAndFoundService {
  lostItems: any[] = JSON.parse(`[
    {
      "id": 1,
      "type": "jewellery",
      "description": "Gold chain with a heart locket",
      "isFound": true,
      "foundId": 1,
      "createdTime": "1648785835432",
      "additionalDetails": null
    }, 
    {
      "id": 2,
      "type": "book",
      "description": "Algorithms and Datastructures",
      "isFound": false,
      "foundId": null,
      "createdTime": "1648785835432",
      "additionalDetails": "CLRS 3d edition"
    }, 
    {
      "id": 3,
      "type": "bag",
      "description": "Laptop bag of DELL",
      "isFound": false,
      "foundId": null,
      "createdTime": "1648785835432",
      "additionalDetails": "Black color with no design"
    }
  ]`);

  foundItems: any[] = JSON.parse(`[
      {
        "id": 1,
        "type": "jewellery",
        "description": "Gold chain with a heart locket",
        "createdTime": "1648785835432",
        "additionalDetails": null
      },
      {
        "id": 2,
        "type": "book",
        "description": "Algorithms and Datastructures",
        "createdTime": "1648785835432",
        "additionalDetails": "CLRS 3d edition"
      },
      {
        "id": 3,
        "type": "bag",
        "description": "Laptop bag of DELL",
        "createdTime": "1648785835432",
        "additionalDetails": "Black color with no design"
      }]`);
  result: any;

  httpOptions = {
    withCredentials: true,
    headers: new HttpHeaders({
      'Content-Type': 'application/json',
      Accept: '*/*',
    })
  };


  isAdmin: boolean = false;

  constructor(private _http: HttpClient) {}

  getLostItems(): Observable<any> {
    return this._http.get(
      'http://localhost:8181/listAllLostItems',
      this.httpOptions
    );
  }

  getFoundItems(): Observable<any> {
  
    return this._http.get(
      'http://localhost:8181/listAllFoundItems',
      this.httpOptions
    );
  }

  createLostRequest(postData : any): Observable<any> {
    return this._http.post(
      'http://localhost:8181/postLostItem',
      postData,
      this.httpOptions
    );
  }

  createFoundRequest(postData : any): Observable<any> {
    return this._http.post(
      'http://localhost:8181/postFoundItem',postData,this.httpOptions
    );
  }

  ngOnInit(): void {
    let isAdmin = localStorage.getItem("isAdmin");
    if (isAdmin) {
      console.log("is admin", isAdmin);
      this.isAdmin = true;
    }
  }


}
