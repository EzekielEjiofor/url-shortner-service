package main

import (
	"log"

	"github.com/ezekielejiofor/url-shortner-service/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func loadRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	app := fiber.New()
	app.Use(logger.New())
	loadRoutes(app)

	log.Fatal(app.Listen(config.AppPort))
}
