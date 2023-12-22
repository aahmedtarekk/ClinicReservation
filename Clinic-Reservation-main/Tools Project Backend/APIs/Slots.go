package APIs

import (
	"toolsProject/controllers"
	"toolsProject/models"

	"github.com/gofiber/fiber/v2"
)

func AddSlot(c *fiber.Ctx) error {
	var slot models.Slot
	err := c.BodyParser(&slot)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	uuid := c.Params("uuid")
	check, slots := controllers.InsertSlot(slot, uuid)
	var response Response
	if !check && len(slots) == 0 {
		response.ResponseStatus = false
		response.ResponseMessage = "User is not signed in"
		return c.Status(200).JSON(response)
	} else if !check {
		response.ResponseStatus = false
		response.ResponseMessage = "Sorry, the inserted Slot conflicts with another slot in the schedule"
		response.UserUUID = uuid
		return c.Status(200).JSON(response)
	}
	response.ResponseStatus = true
	response.ResponseMessage = "Slot was added successfully"
	response.ResponseData = slots
	response.UserUUID = uuid
	return c.Status(200).JSON(response)

}

func CancelSlot(c *fiber.Ctx) error {
	var slot models.Slot
	err := c.BodyParser(&slot)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	uuid := c.Params("uuid")
	var response Response
	if !controllers.CancelSlot(slot, uuid) {
		response.ResponseStatus = false
		response.ResponseMessage = "User is not signed in"
		return c.Status(200).JSON(response)
	}
	response.ResponseStatus = true
	response.ResponseMessage = "Slot was deleted successfully"
	response.UserUUID = uuid
	return c.Status(200).JSON(response)

}

func GetDoctorSlots(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	check, slots := controllers.GetDoctorSlots(uuid)
	var response Response
	if !check {
		response.ResponseStatus = false
		response.ResponseMessage = "User is not signed in"
		return c.Status(200).JSON(response)
	}
	response.ResponseStatus = true
	response.ResponseMessage = "Slots were retrieved successfully"
	response.ResponseData = slots
	response.UserUUID = uuid
	return c.Status(200).JSON(response)
}
