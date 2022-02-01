import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import {RouterModule} from '@angular/router';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';

import {MatToolbarModule} from "@angular/material/toolbar";
import {MatIconModule,  MatButtonModule, MatButtonToggleModule, } from "@angular/material";
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {IvyCarouselModule} from 'angular-responsive-carousel';
import {MatCardModule} from '@angular/material/card';
import { AnnouncementComponent } from './announcement/announcement.component';


@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    AnnouncementComponent
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
    IvyCarouselModule,
    MatCardModule,
    RouterModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
