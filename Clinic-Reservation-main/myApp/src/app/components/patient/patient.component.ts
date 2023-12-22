import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CurrentUser } from 'src/app/CurrentUser';
import { Slot } from '../doctor/slot';
import { MyResponse } from 'src/MyResponse';
import { Appointment } from './Appointment';
import { Doctor } from '../doctor/Doctor';

@Component({
  selector: 'app-patient',
  templateUrl: './patient.component.html',
  styleUrls: ['./patient.component.css']
})
export class PatientComponent implements OnInit{
  constructor(private http : HttpClient, private user: CurrentUser){}
  patientAppointments: Appointment[] = [];
  doctors: Doctor[]
  slots : Slot[]
  times : Set<string> = new Set<string>()
  dates : Set<string> = new Set<string>()
  editMode : boolean = false

  ngOnInit() : void{
    this.http.get<MyResponse>("http://127.0.0.1:3001/ClinicReservation/GetPatientAppointments/"+this.user.uuid).subscribe((response) => {
      if(!response.ResponseStatus){
        alert(response.ResponseMessage)
      }
      else{
        this.patientAppointments = response.ResponseData
      }
    })
    this.http.get<MyResponse>("http://127.0.0.1:3001/ClinicReservation/GetAllDoctors").subscribe((response) => {
      this.doctors = response.ResponseData
    })

  }
  addPatientAppointment() {
    if(!this.editMode){
      var newPatientAppointment = new Appointment()
      this.editMode = true
      this.patientAppointments.push(newPatientAppointment);
      this.resetData()
    }
    else{
      alert("Finish the current appointemnt first !")
    }
  }
  

  editAppointment(appointment: Appointment) {
    const index = this.patientAppointments.indexOf(appointment);
    if (index !== -1 && appointment.DoctorName != "" && appointment.Slot.Date != "" && appointment.Slot.Time != "") {
      var object = {
        "DoctorName" : appointment.DoctorName,
        "Slot" : {
          "Time" : appointment.Slot.Time,
          "Date" : appointment.Slot.Date
        }
      }
      this.http.post("http://127.0.0.1:3001/ClinicReservation/CancelAppointment/"+this.user.uuid, object).subscribe()
      for (let doctor of this.doctors) {
        if(doctor.Name == appointment.DoctorName){
          this.slots = doctor.Slots
          for(let slot of this.slots){
            this.dates.add(slot.Date)
          }
          break
        }
      }
      appointment.Slot.IsReserved = false
      for(let slot of this.slots){
        if (slot.Date == appointment.Slot.Date && slot.Time == appointment.Slot.Time ){
          slot.IsReserved = false
          break;
        }
      }
    }
    this.onDoctorSelection(appointment)
    this.onDateSelection(appointment.Slot.Date)
    this.editMode = appointment.editMode = true;
  }

  saveAppointment(appointment: Appointment) {
    if(appointment.DoctorName == "" || appointment.Slot.Date == "" || appointment.Slot.Time == ""){
      alert("Please finish all the required fields first !")
      return
    }
    this.editMode = appointment.editMode = false;
    var object = {
      "DoctorName" : appointment.DoctorName,
      "Slot" : {
        "Time" : appointment.Slot.Time,
        "Date" : appointment.Slot.Date
      }
    }
    this.http.post<MyResponse>("http://127.0.0.1:3001/ClinicReservation/ReserveAppointment/"+this.user.uuid, object).subscribe((response)=>{
      if(!response.ResponseStatus){
        this.patientAppointments.splice(this.patientAppointments.indexOf(appointment), 1);
        alert(response.ResponseMessage)
      }
    })
    appointment.Slot.IsReserved = true
    for(let slot of this.slots){
      if (slot.Date == appointment.Slot.Date && slot.Time == appointment.Slot.Time ){
        slot.IsReserved = true
        break
      }
    }
    this.resetData()
  }

  cancelAppointment(appointment: Appointment) {
    const index = this.patientAppointments.indexOf(appointment);
    if (index !== -1) {
      this.patientAppointments.splice(index, 1);
      if(appointment.DoctorName != "" && appointment.Slot.Date != "" && appointment.Slot.Time != ""){
        var object = {
          "DoctorName" : appointment.DoctorName,
          "Slot" : {
            "Time" : appointment.Slot.Time,
            "Date" : appointment.Slot.Date
          }
        }
        this.http.post("http://127.0.0.1:3001/ClinicReservation/CancelAppointment/"+this.user.uuid, object).subscribe()
      }
    }
    appointment.Slot.IsReserved = false
    for(let slot of this.slots){
      if (slot.Date == appointment.Slot.Date && slot.Time == appointment.Slot.Time ){
        slot.IsReserved = false
        break;
      }
    }
    this.resetData()
    this.editMode = false
  }
  onDoctorSelection(appointment : Appointment){
    this.resetData()
    for (let doctor of this.doctors) {
      if(doctor.Name == appointment.DoctorName){
        this.slots = doctor.Slots
        for(let slot of this.slots){
          this.dates.add(slot.Date)
        }
        break
      }
    }
    appointment.Slot.Date = ""
    appointment.Slot.Time = ""
  }
  onDateSelection(date : string){
    this.times.clear()
    for(let slot of this.slots){
      if(slot.Date == date){
        this.times.add(slot.Time)
      }
    }
  }
  resetData(){
    this.slots = []
    this.times.clear()
    this.dates.clear()
  }
  isTimeReserved(date : string, time : string) : boolean {
    for(let slot of this.slots){
      if(slot.IsReserved && slot.Date == date && slot.Time == time){
        return true;
      }
    }
    return false
  }
}
