import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { AnnouncementComponent } from './announcement/announcement.component';
import { LostAndFoundComponent } from './lost-and-found/lost-and-found.component';
import { MaintenanceRequestComponent } from './maintenance-request/maintenance-request.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'home',
    pathMatch: 'full',
  },
  { path: 'home', component: HomeComponent },
  { path: 'announcement', component: AnnouncementComponent },
  { path: 'lost-and-found', component: LostAndFoundComponent },
  { path: 'maintenance-request', component: MaintenanceRequestComponent },
];

@NgModule({
  declarations: [],
  imports: [CommonModule, RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
