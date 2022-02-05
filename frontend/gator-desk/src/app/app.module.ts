import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatCardModule } from '@angular/material/card';
import {
  FontAwesomeModule,
  FaIconLibrary,
} from '@fortawesome/angular-fontawesome';
import { fas } from '@fortawesome/free-solid-svg-icons';
import { far } from '@fortawesome/free-regular-svg-icons';
import { fab } from '@fortawesome/free-brands-svg-icons';
import { IvyCarouselModule } from 'angular-responsive-carousel';
import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { HomeComponent } from './home/home.component';
<<<<<<< HEAD
import { FooterComponent } from './shared/footer/footer.component';
import { AnnouncementComponent } from './announcement/announcement.component';
import { HeaderNavComponent } from './shared/header-nav/header-nav.component';
import { HeaderComponent } from './shared/header/header.component';
import { LostAndFoundComponent } from './lost-and-found/lost-and-found.component';
import { MaintenanceRequestComponent } from './maintenance-request/maintenance-request.component';
=======


import {MatToolbarModule} from "@angular/material/toolbar";

import {MatIconModule,  MatButtonModule, MatButtonToggleModule, MatFormFieldModule, MatSelectModule, MatOptionModule,
   MatInputModule , MatExpansionModule} from "@angular/material";

import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatCardModule} from '@angular/material/card';
import { FooterComponent } from './footer/footer.component';
import { AnnouncementComponent } from './announcement/announcement.component';
import { AdminAnnouncemetComponent } from './admin-announcemet/admin-announcemet.component';
import { IvyCarouselModule } from 'angular-responsive-carousel';


>>>>>>> 449d1936af81858ad52be6f93eef0b5016a31890

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    FooterComponent,
    AnnouncementComponent,
<<<<<<< HEAD
    HeaderNavComponent,
    HeaderComponent,
    LostAndFoundComponent,
    MaintenanceRequestComponent,
=======
    AdminAnnouncemetComponent
>>>>>>> 449d1936af81858ad52be6f93eef0b5016a31890
  ],
  imports: [
    BrowserModule,
    MatGridListModule,
    MatIconModule,
    MatToolbarModule,
    MatButtonModule,
<<<<<<< HEAD
=======
    MatButtonToggleModule,
    BrowserModule,
    BrowserAnimationsModule,
>>>>>>> 449d1936af81858ad52be6f93eef0b5016a31890
    MatCardModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    RouterModule,
<<<<<<< HEAD
    FontAwesomeModule,
    IvyCarouselModule,
=======
    MatFormFieldModule,
    IvyCarouselModule,
    MatSelectModule,
    MatOptionModule,
    MatInputModule,
    MatExpansionModule
>>>>>>> 449d1936af81858ad52be6f93eef0b5016a31890
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {
  constructor(library: FaIconLibrary) {
    library.addIconPacks(fas, far, fab);
  }
}
