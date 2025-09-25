package main

import (
	dal "fiber/DAL"
	"fiber/database"
	"fiber/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&dal.Todo{})
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	type TodoCreate struct {
		Title string
	}

	app.Post("/todos", services.CreateTodo)
	app.Get("/todos", services.GetTodos)
	app.Get("/todos/:todoID", services.GetTodo)
	app.Listen("localhost:3000")
}
