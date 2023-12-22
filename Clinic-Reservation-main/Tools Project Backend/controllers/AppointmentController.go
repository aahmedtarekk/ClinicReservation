package controllers

import (
	"toolsProject/initializers"
	"toolsProject/models"
)

func ReserveAppointment(appointment models.Appointment, uuid string) (bool, models.Appointment) {
	if userPair, check := initializers.ActivePool[uuid]; !check {
		return false, models.Appointment{}
	} else {
		appointment.PatientID = userPair.ID
	}
	DoctorID := GetDoctorID(appointment.DoctorName)
	slot := GetSlot(appointment.Slot.Date, appointment.Slot.Time, DoctorID)
	if slot.IsReserved {
		return false, appointment
	}
	UpdateSlotStatus(&slot, true)
	appointment.SlotID = slot.ID
	appointment.Slot = slot
	AddAppointment(&appointment)
	return true, appointment
}

func CancelAppointment(appointment models.Appointment, uuid string) (bool, []models.Appointment) {
	if userPair, check := initializers.ActivePool[uuid]; !check {
		return false, []models.Appointment{}
	} else {
		appointment.PatientID = userPair.ID
	}
	DoctorID := GetDoctorID(appointment.DoctorName)
	slot := GetSlot(appointment.Slot.Date, appointment.Slot.Time, DoctorID)
	DeleteAppointment(slot.ID, appointment.DoctorName)
	UpdateSlotStatus(&slot, false)
	return true, GetAllAppointments(appointment.PatientID)
}

func ChangeAppointment(oldAppointment models.Appointment, newAppointment models.Appointment, uuid string) (bool, models.Appointment) {
	if _, check := initializers.ActivePool[uuid]; !check {
		return false, models.Appointment{}
	}
	DoctorID := GetDoctorID(newAppointment.DoctorName)
	slot := GetSlot(newAppointment.Slot.Date, newAppointment.Slot.Time, DoctorID)
	newAppointment.SlotID = slot.ID
	if slot.IsReserved {
		return false, oldAppointment
	}
	UpdateSlotStatus(&slot, true)
	DoctorID = GetDoctorID(oldAppointment.DoctorName)
	slot = GetSlot(oldAppointment.Slot.Date, oldAppointment.Slot.Time, DoctorID)
	UpdateSlotStatus(&slot, false)
	oldAppointment.SlotID = slot.ID
	GetAppointment(&oldAppointment)
	UpdateAppointment(&oldAppointment, newAppointment)
	return true, oldAppointment
}

func GetPatientAppointments(uuid string) (bool, []models.Appointment) {
	var id uint
	if userPair, check := initializers.ActivePool[uuid]; !check {
		return false, []models.Appointment{}
	} else {
		id = userPair.ID
	}
	return true, GetAllAppointments(id)
}
