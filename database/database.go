package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("myDatabase.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Database Connection Error", err)
	}
	DB = db
	fmt.Println("Connected To Database")
}
