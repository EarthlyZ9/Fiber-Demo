package routes

import (
	"github.com/EarthlyZ9/fiberdemo/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(route fiber.Router) {
	route.Get("", handlers.GetUsers)
	route.Get("/:id", handlers.GetUser)
	route.Post("", handlers.CreateUser)
	route.Delete("/:id", handlers.DeleteUser)
	route.Patch("/:id", handlers.UpdateUser)
}
