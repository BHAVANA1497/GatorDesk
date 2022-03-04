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
import { FooterComponent } from './shared/footer/footer.component';
import { AnnouncementComponent } from './announcement/announcement.component';
import { HeaderNavComponent } from './shared/header-nav/header-nav.component';
import { LostAndFoundComponent } from './lost-and-found/lost-and-found.component';
import { MaintenanceRequestComponent } from './maintenance-request/maintenance-request.component';

import { MatButtonToggleModule} from '@angular/material/button-toggle';
import{MatFormFieldModule} from'@angular/material/form-field';
import { MatSelectModule} from '@angular/material/select';
import {    MatInputModule} from '@angular/material/input';
import { MatExpansionModule } from '@angular/material/expansion';

import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import { MatOptionModule } from '@angular/material/core';
import { AdminAnnouncementComponent } from './admin-announcemet/admin-announcement.component';
import { AdminAnnouncementService } from './admin-announcemet/admin-annoucement.service';
import { CommonService } from './providers/common.service';
import { UserService } from './providers/user.service';
import { LoginComponent } from './login/login.component';
import { HttpClientModule } from '@angular/common/http';
import { AnnouncemetService } from './announcement/annoucement.service';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    FooterComponent,
    AnnouncementComponent,
    HeaderNavComponent,
    LostAndFoundComponent,
    MaintenanceRequestComponent,
    AdminAnnouncementComponent,
    LoginComponent,
  ],
  imports: [
    BrowserModule,
    MatGridListModule,
    MatIconModule,
    MatToolbarModule,
    MatButtonModule,
    MatCardModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    RouterModule,
    FontAwesomeModule,
    IvyCarouselModule,
    MatSelectModule,
    MatInputModule,
    MatExpansionModule,
    MatButtonToggleModule,
    MatFormFieldModule,
    FormsModule,
    ReactiveFormsModule,
    MatIconModule,
    MatOptionModule,
    HttpClientModule,
  ],
  providers: [
    AdminAnnouncementService,
    AnnouncemetService,
    CommonService,
    UserService
  ],
  bootstrap: [AppComponent],
})
export class AppModule {
  constructor(library: FaIconLibrary) {
    library.addIconPacks(fas, far, fab);
  }
}
