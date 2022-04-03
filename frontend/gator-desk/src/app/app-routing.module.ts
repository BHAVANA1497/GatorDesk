import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { AnnouncementComponent } from './announcement/announcement.component';
import { LostAndFoundComponent } from './lost-and-found/lost-and-found.component';
import { MaintenanceRequestComponent } from './maintenance-request/maintenance-request.component';
import { AdminAnnouncementComponent } from './admin-announcemet/admin-announcement.component';
import { SignupComponent } from './signup/signup.component';
import { LoginComponent } from './login/login.component';
import { AdminLostAndFoundComponent } from './admin-lost-and-found/admin-lost-and-found.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'home',
    pathMatch: 'full',
  },
  { path: 'home', component: HomeComponent },
  { path: 'announcement', component: AnnouncementComponent },
  { path: 'lost-and-found', component: AdminLostAndFoundComponent },
  { path: 'maintenance-request', component: LostAndFoundComponent },
  { path: 'login', component: LoginComponent },
  { path : 'signup', component: SignupComponent}
];

@NgModule({
  declarations: [],
  imports: [CommonModule, RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
