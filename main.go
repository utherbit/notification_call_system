package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"notificator/connections"
	"notificator/handlers"
	"notificator/utilities"
)

func main() {
	utilities.CheckEnvFile()
	connections.InitPostgresConnection()
	connections.InitAsterisk()
	connections.InitSynthesizerApi()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Post("/add_call", handlers.HandlerNewCall)
	app.Post("/check_status", handlers.HandlerCheckStatusCall)
	app.Post("/start_call", handlers.HandlerStartCall)
	app.Post("/asterisk_result", handlers.HandlerAsteriskResult)

	log.Fatal(app.Listen("0.0.0.0:4200"))
}
