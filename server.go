package main

import (
	"fmt"
	"log"

	configuration "github.com/EarthlyZ9/fiberdemo/config"
	"github.com/EarthlyZ9/fiberdemo/database"
	"github.com/EarthlyZ9/fiberdemo/models"
	"github.com/EarthlyZ9/fiberdemo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func setupRoutes(app *fiber.App) {
	// Routes
	v1 := app.Group("/api/v1")

	// Users
	userRoutes := v1.Group("/users")
	routes.UserRoutes(userRoutes)

	// Posts
	postRoutes := v1.Group("/posts")
	routes.PostRoutes(postRoutes)
}

func main() {
	config := configuration.New()

	app := fiber.New(*config.GetFiberConfig())

	app.Use(recover.New())

	// Initialize database
	db, err := database.New(&database.DatabaseConfig{
		Driver:   config.GetString("DB_DRIVER"),
		Host:     config.GetString("DB_HOST"),
		Username: config.GetString("DB_USERNAME"),
		Password: config.GetString("DB_PASSWORD"),
		Port:     config.GetInt("DB_PORT"),
		Database: config.GetString("DB_DATABASE"),
	})

	// Auto-migrate database models
	if err != nil {
		fmt.Println("failed to connect to database:", err.Error())
	} else {
		if db == nil {
			fmt.Println("failed to connect to database: db variable is nil")
		} else {
			err := db.AutoMigrate(&models.User{})
			if err != nil {
				fmt.Println("failed to automigrate user model:", err.Error())
				return
			}
			err = db.AutoMigrate(&models.Post{})
			if err != nil {
				fmt.Println("failed to automigrate post model:", err.Error())
				return
			}
		}
	}

	// Routes
	setupRoutes(app)

	// Root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
