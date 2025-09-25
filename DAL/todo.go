package dal

import (
	"fiber/database"

	"gorm.io/gorm"
)

func CreateTodo(todo *Todo) *gorm.DB {
	return database.DB.Create(&todo)
}

func GetTodos(dest any) *gorm.DB {
	return database.DB.Model(&Todo{}).Find(dest)
}

type Todo struct {
	ID        int
	Title     string
	Completed bool `gorm:"default:false"`
}

func GetTodoByID(dest any, id any) *gorm.DB {
	return database.DB.Model(&Todo{}).Where("id = ?", id).First(dest)
}
