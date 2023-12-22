export class Slot{
    ID : number;
    Date : string;
    Time : string;
    DoctorID : number;
    IsReserved : boolean;
    editMode?: boolean;
    constructor(){
        this.ID = 0
        this.Date = ""
        this.Time = ""
        this.DoctorID = 0
        this.IsReserved = false
        this.editMode = true
    }
}