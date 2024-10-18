package main 

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"go-fiber-websockets/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/api/message", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, Restful API"})
	})

	app.Get("/ws", websocket.New(handlers.WebsocketHandler))


	log.Fatal(app.Listen(":8080"))
}