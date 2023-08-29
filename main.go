package main

import (
	"go-todo-api/app/routes"
	"go-todo-api/app/dal"
	"go-todo-api/database"

	"os"
	_ "github.com/joho/godotenv/autoload"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return ":" + port
}

func main() {
	database.ConnectDb()
	database.Migrate(&dal.User{}, &dal.Todo{})
	app := fiber.New()

	routes.AuthRoutes(app)
	routes.TodoRoutes(app)

	app.Use(cors.New(cors.Config{
		//AllowOrigins: "*",
		//AllowHeaders: "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(getPort())
}
