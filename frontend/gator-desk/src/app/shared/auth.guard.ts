import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { AuthenticationService } from '../providers/authentication.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(private service : AuthenticationService){

  }
  canActivate() {
    if(this.service.isLoggedIn()){
      return true;
    }else{
      return false;
    }
  }
  
}
