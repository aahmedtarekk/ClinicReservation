package APIs

import (
	"toolsProject/controllers"
	"toolsProject/models"

	"github.com/gofiber/fiber/v2"
)

func ReserveAppointment(c *fiber.Ctx) error {
	var appointment models.Appointment
	err := c.BodyParser(&appointment)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var check bool
	uuid := c.Params("uuid")
	check, appointment = controllers.ReserveAppointment(appointment, uuid)
	var response Response
	if !check && appointment.DoctorName == "" {
		response.ResponseStatus = false
		response.ResponseMessage = "User not signed in !"
		return c.Status(200).JSON(response)
	} else if !check {
		response.ResponseStatus = false
		response.ResponseMessage = "Sorry, Slot is already taken"
		response.UserUUID = uuid
		return c.Status(200).JSON(response)
	}
	response.ResponseStatus = true
	response.ResponseMessage = "Appointment was reserved successfully"
	response.ResponseData = appointment
	response.UserUUID = uuid
	return c.Status(200).JSON(response)
}

func CancelAppointment(c *fiber.Ctx) error {
	var appointment models.Appointment
	err := c.BodyParser(&appointment)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	uuid := c.Params("uuid")
	check, appointments := controllers.CancelAppointment(appointment, uuid)
	var response Response
	if !check {
		response.ResponseStatus = false
		response.ResponseMessage = "User not signed in !"
		return c.Status(200).JSON(response)
	}
	response.ResponseStatus = true
	response.ResponseMessage = "Appointment was canceled successfully"
	response.ResponseData = appointments
	response.UserUUID = uuid
	return c.Status(200).JSON(response)
}

func ChangeAppointment(c *fiber.Ctx) error {
	var appointmentsPair models.AppointmentsPair
	err := c.BodyParser(&appointmentsPair)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	uuid := c.Params("uuid")
	check, newAppointment := controllers.ChangeAppointment(appointmentsPair.OldAppointment, appointmentsPair.NewAppointment, uuid)
	var response Response
	if !check && newAppointment.DoctorName == "" {
		response.ResponseStatus = false
		response.ResponseMessage = "User not signed in !"
		return c.Status(200).JSON(response)
	} else if !check {
		response.ResponseStatus = false
		response.ResponseMessage = "Sorry, Slot is already taken"
		response.UserUUID = uuid
		return c.Status(200).JSON(response)
	}
	response.ResponseStatus = true
	response.ResponseMessage = "Appointment was changed successfully"
	response.ResponseData = newAppointment
	response.UserUUID = uuid
	return c.Status(200).JSON(response)
}

func GetPatientAppointments(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	check, appointments := controllers.GetPatientAppointments(uuid)
	var response Response
	if !check {
		response.ResponseStatus = false
		response.ResponseMessage = "User not Signed in !"
		return c.Status(200).JSON(response)
	}
	response.ResponseStatus = true
	response.ResponseMessage = "Appointments were retrieved successfully"
	response.ResponseData = appointments
	response.UserUUID = uuid
	return c.Status(200).JSON(response)
}
