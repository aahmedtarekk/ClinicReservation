import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SignInComponent } from './components/sign-in/sign-in.component';
import { SignUpComponent } from './components/sign-up/sign-up.component';
import { DoctorComponent } from './components/doctor/doctor.component';
import { PatientComponent } from './components/patient/patient.component';
const routes: Routes = [
  {path:'sign-in',component:SignInComponent},
  {path:'sign-up',component:SignUpComponent},
  {path:'doctor',component:DoctorComponent},
  {path:'patient',component:PatientComponent}

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
