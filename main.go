package main

import (
	"flag"
	"gofiber-todo/config/database"
	"gofiber-todo/src/entity"
	auth_module "gofiber-todo/src/modules/auth"
	todo_module "gofiber-todo/src/modules/todo"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	database.ConnectDb()
	database.Migrate(&entity.Todo{}, &entity.User{})

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		// Format: "${pid} ${status} - ${method} ${path}\n",
	}))

	// Create a /api endpoint
	apiRoute := app.Group("/api")

	todo_module.NewModule(apiRoute)
	auth_module.NewModule(apiRoute)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
