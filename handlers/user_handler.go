package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	return c.SendString("Get all users")
}

func GetUser(c *fiber.Ctx) error {
	return c.SendString("Get user by id")
}

func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Update an user")
}

func CreateUser(c *fiber.Ctx) error {
	return c.SendString("Create user")
}

func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Delete an user")
}
