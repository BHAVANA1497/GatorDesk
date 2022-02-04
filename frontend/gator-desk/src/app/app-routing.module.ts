import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AnnouncementComponent } from './announcement/announcement.component';
import { HomeComponent } from './home/home.component';
import { AdminAnnouncemetComponent } from './admin-announcemet/admin-announcemet.component';


const routes: Routes = [
    { path: '', component: HomeComponent},
    { path: 'home', component: HomeComponent },
    { path: 'announcement', component: AnnouncementComponent },
    {path: 'login', component: AdminAnnouncemetComponent}

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
