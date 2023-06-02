package routes

import (
	"github.com/EarthlyZ9/fiberdemo/handlers"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(route fiber.Router) {
	route.Get("", handlers.GetPosts)
	route.Get("/:id", handlers.GetPosts)
	route.Post("", handlers.CreatePost)
	route.Delete("/:id", handlers.DeletePost)
	route.Patch("/:id", handlers.UpdatePost)
}
