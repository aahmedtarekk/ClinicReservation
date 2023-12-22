package APIs

import (
	"toolsProject/controllers"

	"github.com/gofiber/fiber/v2"
)

func GetAllDoctors(c *fiber.Ctx) error {
	var response Response
	response.ResponseData = controllers.GetAllDoctors()
	response.ResponseStatus = true
	response.ResponseMessage = "Doctors were retrieved successfully"
	return c.Status(200).JSON(response)
}
