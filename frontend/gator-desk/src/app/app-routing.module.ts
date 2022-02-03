import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AnnouncementComponent } from './announcement/announcement.component';
import { HomeComponent } from './home/home.component';
import { FooterComponent } from './footer/footer.component';


const routes: Routes = [
    { path: '', component: HomeComponent,
      children:[
        {
          path:'',
          component: FooterComponent
        }
      ]
  },
    { path: 'home', component: HomeComponent },
    { path: 'announcement', component: AnnouncementComponent },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
