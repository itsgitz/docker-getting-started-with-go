package app

import (
	"log"
	"music-playlist/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() {
	// Set fiber application and enable recover middleware
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "AppGitzFiber",
	})
	app.Use(recover.New())

	// Define routers
	router.New(app)

	// Run the server
	log.Fatal(app.Listen(":3000"))
}
