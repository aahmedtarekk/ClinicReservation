package models

type Patient struct {
	ID           uint `gorm:primaryKey`
	Name         string
	Mail         string
	Password     string
	Type         string
	Appointments []Appointment
}
