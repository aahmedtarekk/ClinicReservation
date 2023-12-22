package controllers

import (
	"toolsProject/initializers"
	"toolsProject/models"
)

func AddDoctor(doctor models.Doctor) models.Doctor {
	initializers.DB.Create(&doctor)
	return doctor
}

func AddPatient(patient models.Patient) models.Patient {
	initializers.DB.Create(&patient)
	return patient
}

func GetDoctorID(name string) uint {
	var doctor models.Doctor
	initializers.DB.Find(&doctor, "name = ?", name)
	return doctor.ID
}

func GetAllDoctors() []models.Doctor {
	doctors := []models.Doctor{}
	initializers.DB.Preload("Slots").Find(&doctors)
	return doctors
}

func GetAllPatients() []models.Patient {
	patient := []models.Patient{}
	initializers.DB.Find(&patient)
	return patient
}

func AddSlot(slot *models.Slot) {
	initializers.DB.Create(&slot)
}

func UpdateSlotStatus(slot *models.Slot, status bool) {
	initializers.DB.Model(&slot).Update("IsReserved", status)
}

func GetSlot(Date string, Time string, DoctorID uint) models.Slot {
	var slot models.Slot
	initializers.DB.Find(&slot, "doctor_id = ? AND Date = ? AND Time = ?", DoctorID, Date, Time)
	return slot
}

func DeleteSlot(slot models.Slot, DoctorID uint) {
	initializers.DB.Delete(&models.Slot{}, "Date = ? AND Time = ? And doctor_id = ?", slot.Date, slot.Time, DoctorID)
}

func GetAllDoctorSlots(id uint) []models.Slot {
	var slots []models.Slot
	initializers.DB.Find(&slots, "doctor_id = ?", id)
	return slots
}

func GetAppointment(appointment *models.Appointment) {
	initializers.DB.Preload("Slot").Find(&appointment, "slot_id = ?", appointment.SlotID)
}

func GetAllAppointments(PatientID uint) []models.Appointment {
	var appointments []models.Appointment
	initializers.DB.Preload("Slot").Find(&appointments, "patient_id = ?", PatientID)
	return appointments
}

func AddAppointment(appointment *models.Appointment) {
	initializers.DB.Create(&appointment)
}

func DeleteAppointment(SlotID uint, DoctorName string) {
	initializers.DB.Delete(&models.Appointment{}, "doctor_name = ? AND slot_id = ?", DoctorName, SlotID)
}

func UpdateAppointment(oldAppointment *models.Appointment, newAppointment models.Appointment) {
	initializers.DB.Model(&oldAppointment).Updates(models.Appointment{SlotID: newAppointment.SlotID, DoctorName: newAppointment.DoctorName})
	initializers.DB.Preload("Slot").Find(&oldAppointment)
}
