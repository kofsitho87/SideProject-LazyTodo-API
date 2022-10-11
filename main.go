package main

import (
	"flag"
	"gofiber-todo/src/config"
	"gofiber-todo/src/config/database"
	"gofiber-todo/src/entity"
	auth_module "gofiber-todo/src/modules/auth"
	todo_module "gofiber-todo/src/modules/todo"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	_ "gofiber-todo/docs"
)

// @title Todo API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:3000
// @BasePath /
func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	database.ConnectDb()
	database.Migrate(&entity.Todo{}, &entity.User{})

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: false, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		// Format: "${pid} ${status} - ${method} ${path}\n",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// Create a /api endpoint
	apiRoute := app.Group("/api")

	todo_module.NewModule(apiRoute)
	auth_module.NewModule(apiRoute)

	// Listen on port 3000
	log.Fatal(app.Listen(":" + config.PORT)) // go run app.go -port=:3000
}
