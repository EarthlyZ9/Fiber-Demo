package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	return c.SendString("Get all posts of a post")
}

func GetPost(c *fiber.Ctx) error {
	return c.SendString("Get post by id")
}

func UpdatePost(c *fiber.Ctx) error {
	return c.SendString("Update a post")
}

func CreatePost(c *fiber.Ctx) error {
	return c.SendString("Create post")
}

func DeletePost(c *fiber.Ctx) error {
	return c.SendString("Delete a post")
}
