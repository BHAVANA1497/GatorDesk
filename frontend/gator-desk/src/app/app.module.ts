import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import {RouterModule} from '@angular/router';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';

import {MatToolbarModule} from "@angular/material/toolbar";
import {MatIconModule,  MatButtonModule, MatButtonToggleModule, MatFormFieldModule, MatSelectModule, MatOptionModule, MatInputModule } from "@angular/material";
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatCardModule} from '@angular/material/card';
import { FooterComponent } from './footer/footer.component';
import { AnnouncementComponent } from './announcement/announcement.component';
import { AdminAnnouncemetComponent } from './admin-announcemet/admin-announcemet.component';
import { IvyCarouselModule } from 'angular-responsive-carousel';


@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    FooterComponent,
    AnnouncementComponent,
    AdminAnnouncemetComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MatToolbarModule,
    MatIconModule,
    MatButtonModule,
    MatButtonToggleModule,
    BrowserModule,
    BrowserAnimationsModule,
    MatCardModule,
    RouterModule,
    MatFormFieldModule,
    IvyCarouselModule,
    MatSelectModule,
    MatOptionModule,
    MatInputModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
