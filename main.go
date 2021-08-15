package main

import (
	"fmt"

	"github.com/fatopato/go-fiber-assistant/database"
	"github.com/fatopato/go-fiber-assistant/todo"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/todo", todo.GetAllTODOs)
	v1.Post("/todo", todo.SaveTODO)
	v1.Get("/todo/:id", todo.GetTODOById)
	v1.Delete("/todo/:id", todo.DeleteTODOById)
	v1.Get("/todo/:id/over-time", todo.CheckOverTimeById)
}

func initDB() {
	var err error
	database.DB, err = gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic("DB connection error!!!")
	}
	fmt.Println("DB connection is established")
	database.DB.AutoMigrate(&todo.Todo{})
	fmt.Println("Automigration is completed")
}

func main() {
	app := fiber.New()
	initDB()
	defer database.DB.Close()

	setupRoutes(app)
	app.Listen(3000)
}
