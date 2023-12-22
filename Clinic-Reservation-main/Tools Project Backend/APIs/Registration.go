package APIs

import (
	"toolsProject/controllers"
	"toolsProject/models"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	user = controllers.SignUp(user)
	var response Response
	if user.ID == 0 {
		response.ResponseStatus = false
		response.ResponseMessage = "Name or mail is already taken"
		return c.Status(200).JSON(response)
	}
	response.ResponseStatus = true
	response.ResponseMessage = "User singed up successfully"
	response.ResponseData = user
	response.UserUUID = user.UUID
	return c.Status(200).JSON(response)
}

func SignIn(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var response Response
	user = controllers.SignIn(user)
	if user.ID == 0 {
		response.ResponseStatus = false
		response.ResponseMessage = "User is not signed in"
		return c.Status(200).JSON(response)
	}
	response.ResponseStatus = true
	response.ResponseMessage = "User signed in successfully"
	response.ResponseData = user
	response.UserUUID = user.UUID
	return c.Status(200).JSON(response)
}
