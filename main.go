package main

import (
	"fiber-file-upload/config"
	"fiber-file-upload/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitCloudinary()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("api is ready.....")
	})

	app.Post("/upload", handlers.UploadSingle)
	app.Post("/uploads", handlers.UploadMultiple)

	log.Fatal(app.Listen(":8000"))
}
