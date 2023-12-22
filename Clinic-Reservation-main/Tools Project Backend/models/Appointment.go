package models

type Appointment struct {
	ID         uint `gorm:primaryKey`
	DoctorName string
	SlotID     uint
	Slot       Slot
	PatientID  uint
}
