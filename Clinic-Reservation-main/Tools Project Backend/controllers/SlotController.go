package controllers

import (
	"toolsProject/initializers"
	"toolsProject/models"
)

func InsertSlot(slot models.Slot, uuid string) (bool, []models.Slot) {
	var slots []models.Slot
	var id uint
	if userPair, check := initializers.ActivePool[uuid]; !check {
		return false, []models.Slot{}
	} else {
		id = userPair.ID
	}
	slots = GetAllDoctorSlots(id)
	check := true
	for _, s := range slots {
		if s.Date == slot.Date && s.Time == slot.Time {
			check = false
			break
		}
	}
	if !check {
		return false, slots
	}
	slot.DoctorID = id
	slot.IsReserved = false
	AddSlot(&slot)
	slots = append(slots, slot)
	return true, slots
}

func CancelSlot(slot models.Slot, uuid string) bool {
	var DoctorID uint
	if userPair, check := initializers.ActivePool[uuid]; !check {
		return false
	} else {
		DoctorID = userPair.ID
	}
	DeleteSlot(slot, DoctorID)
	return true
}

func GetDoctorSlots(uuid string) (bool, []models.Slot) {
	var id uint
	if userPair, check := initializers.ActivePool[uuid]; !check {
		return false, []models.Slot{}
	} else {
		id = userPair.ID
	}
	return true, GetAllDoctorSlots(id)
}
