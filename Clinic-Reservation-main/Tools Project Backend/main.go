package main

import (
	"toolsProject/APIs"
	"toolsProject/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func setAPIs(app *fiber.App) {
	app.Post("/ClinicReservation/SignUp", APIs.SignUp)
	app.Post("/ClinicReservation/SignIn", APIs.SignIn)
	app.Post("/ClinicReservation/AddSlot/:uuid", APIs.AddSlot)
	app.Post("/ClinicReservation/ReserveAppointment/:uuid", APIs.ReserveAppointment)
	app.Put("/ClinicReservation/ChangeAppointment/:uuid", APIs.ChangeAppointment)
	app.Post("/ClinicReservation/CancelAppointment/:uuid", APIs.CancelAppointment)
	app.Get("/ClinicReservation/GetPatientAppointments/:uuid", APIs.GetPatientAppointments)
	app.Get("/ClinicReservation/GetDoctorSlots/:uuid", APIs.GetDoctorSlots)
	app.Get("/ClinicReservation/GetAllDoctors", APIs.GetAllDoctors)
	app.Post("/ClinicReservation/CancelSlot/:uuid", APIs.CancelSlot)
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	setAPIs(app)
	app.Listen(":41750")
}
