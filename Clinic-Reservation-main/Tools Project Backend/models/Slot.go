package models

type Slot struct {
	ID         uint `gorm:primaryKey`
	Date       string
	Time       string
	DoctorID   uint
	IsReserved bool
}
