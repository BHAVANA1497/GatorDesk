import { Injectable } from '@angular/core';

@Injectable()
export class UserService {

  _details : any ;
  _token = '';

  constructor() {
    this._token = localStorage.getItem('TOKEN') || '';
    //this._details = JSON.parse(localStorage.getItem('USER_DETAILS')) || null;
  }

}
