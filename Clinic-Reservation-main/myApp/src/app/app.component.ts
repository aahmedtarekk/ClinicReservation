// app.component.ts

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

interface Appointment {
  date: string;
  time: string;
  editMode?: boolean;
}

interface PatientAppointment {
  date: string;
  time: string;
  doctorName: string;
  editMode?: boolean;
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  title = 'myApp';

  doctorAppointments: Appointment[] = [
    { date: '2023-11-07', time: '10:00 AM' },
  ];

  patientAppointments: PatientAppointment[] = [
    { date: '2023-11-08', time: '02:30 PM', doctorName: 'Dr. Smith' },
  ];

  constructor(private router: Router) {}

  ngOnInit() {
    // Automatically navigate to the SignInComponent
    this.router.navigate(['/sign-in']);
  }

  addAppointment() {
    const newAppointment: Appointment = { date: '', time: '', editMode: true };
    this.doctorAppointments.push(newAppointment);
  }

  editAppointment(appointment: Appointment) {
    appointment.editMode = true;
  }

  saveAppointment(appointment: Appointment) {
    appointment.editMode = false;
    console.log('Save appointment:', appointment);
  }

  cancelAppointment(appointment: Appointment | PatientAppointment) {
    if ('doctorName' in appointment) {
      const index = this.patientAppointments.indexOf(appointment as PatientAppointment);
      if (index !== -1) {
        this.patientAppointments.splice(index, 1);
      }
    } else {
      const index = this.doctorAppointments.indexOf(appointment as Appointment);
      if (index !== -1) {
        this.doctorAppointments.splice(index, 1);
      }
    }
  }

  addPatientAppointment() {
    const newPatientAppointment: PatientAppointment = { date: '', time: '', doctorName: '', editMode: true };
    this.patientAppointments.push(newPatientAppointment);
  }
}
