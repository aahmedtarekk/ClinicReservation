import { Slot } from "../doctor/slot";

export class Appointment{
    ID: number;
    DoctorName: string;
    SlotID: number;
    Slot : Slot;
    PatientID: number;
    editMode?: boolean;
    constructor(){
        this.ID = 0
        this.DoctorName = ""
        this.SlotID = 0
        this.Slot = new Slot()
        this.PatientID = 0
        this.editMode = true
    }
}