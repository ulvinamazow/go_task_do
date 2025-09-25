package services

import (
	"errors"
	dal "fiber/DAL"
	"fiber/types"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateTodo(c *fiber.Ctx) error {
	t := new(types.TodoCreateDTO)
	err := c.BodyParser(t)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	if t.Title == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Title can not be empty",
		})
	}
	newTodo := dal.Todo{
		Title: t.Title,
	}
	res := dal.CreateTodo(&newTodo)
	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to Creat Todo",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Todo Created Successfully.",
	})
}

func GetTodos(c *fiber.Ctx) error {
	todos := []types.TodoResponse{}
	res := dal.GetTodos(&todos)
	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to Creat Todo",
		})
	}
	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID")
	d := types.TodoResponse{}
	res := dal.GetTodoByID(&d, todoID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "No Todo Found Matching The Id Number",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed To Get Todo",
		})
	}
	return c.JSON(d)
}
