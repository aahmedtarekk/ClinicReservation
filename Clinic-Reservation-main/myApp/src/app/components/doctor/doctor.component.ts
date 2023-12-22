import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CurrentUser } from 'src/app/CurrentUser';
import { MyResponse } from 'src/MyResponse';
import { Slot } from './slot';

@Component({
  selector: 'app-doctor',
  templateUrl: './doctor.component.html',
  styleUrls: ['./doctor.component.css']
})
export class DoctorComponent implements OnInit{

  constructor(private user: CurrentUser, private http : HttpClient){}

  doctorSlots: Slot[] = [];

  ngOnInit() : void{
    this.http.get<MyResponse>("http://127.0.0.1:3001/ClinicReservation/GetDoctorSlots/"+this.user.uuid).subscribe((response) => {
      this.doctorSlots = response.ResponseData
    })
  }

  addSlot() {
    var newSlot: Slot = new Slot()
    this.doctorSlots.push(newSlot);
  }

  editSlot(slot: Slot) {
    if(slot.Date != ""){
      var object = {
        "Date" : slot.Date,
        "Time" : slot.Time
      }
      this.http.post<MyResponse>("http://127.0.0.1:3001/ClinicReservation/CancelSlot/"+this.user.uuid, object).subscribe()
    }
    slot.editMode = true;
  }

  saveSlot(slot: Slot) {
    slot.editMode = false;
    var object = {
      "Time" : slot.Time,
      "Date" : slot.Date
    }
    this.http.post<MyResponse>("http://127.0.0.1:3001/ClinicReservation/AddSlot/"+this.user.uuid, object).subscribe((response) => {
      if(!response.ResponseStatus){
        this.doctorSlots.splice(this.doctorSlots.indexOf(slot), 1);
        alert(response.ResponseMessage)
      }
    })
  }

  cancelSlot(slot: Slot) {
    const index = this.doctorSlots.indexOf(slot);
    if (index !== -1) {
      var object = {
        "Date" : slot.Date,
        "Time" : slot.Time
      }
      if(slot.Date != ""){
        this.http.post<MyResponse>("http://127.0.0.1:3001/ClinicReservation/CancelSlot/"+this.user.uuid, object).subscribe()
      }
      this.doctorSlots.splice(index, 1);
    }
  }
}
