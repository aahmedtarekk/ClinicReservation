import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SignInComponent } from './components/sign-in/sign-in.component';
import { SignUpComponent } from './components/sign-up/sign-up.component';
import { DoctorComponent } from './components/doctor/doctor.component';
import { PatientComponent } from './components/patient/patient.component';
import {HttpClientModule} from '@angular/common/http';
import { CurrentUser } from './CurrentUser';

@NgModule({
  declarations: [
    AppComponent,
    SignInComponent,
    SignUpComponent,
    DoctorComponent,
    PatientComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [CurrentUser],
  bootstrap: [AppComponent]
})
export class AppModule { }
