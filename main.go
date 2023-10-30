package main

import (
	"github.com/gofiber/fiber/v2"
    "github.com/dpi0/fiber_api0/api/routes"
)

func main() {
    app := fiber.New()

    // app.Get("/", func(c *fiber.Ctx) error {
    //     return c.SendString("Hello, World!")
    // })

	routes.SetupRoutes(app)

    app.Listen(":9001")
}

