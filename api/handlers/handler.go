package handlers

import (
    "github.com/gofiber/fiber/v2"
    // "github.com/dpi0/fiber_api0/api/models"
)

// Handler to get all users
// func GetUsers(c *fiber.Ctx) error {
//     // Implement logic to fetch and return all users
//     return c.Status(fiber.StatusOK).JSON(users)
// }

// Other handler functions for creating, updating, or deleting users

func Root(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}