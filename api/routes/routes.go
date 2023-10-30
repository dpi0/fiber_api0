package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/dpi0/fiber_api0/api/handlers"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")

	
    api.Get("/", handlers.Root)

    // User endpoints
    // api.Get("/users", handlers.GetUsers)
    // Other CRUD operations for users
}
