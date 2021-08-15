package todo

import (
	"time"

	"github.com/fatopato/go-fiber-assistant/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title     string    `json:title`
	Completed bool      `json:completed`
	DueTime   time.Time `json:dueTime`
}

func GetAllTODOs(c *fiber.Ctx) {
	db := database.DB
	var todos []Todo
	db.Find(&todos)
	c.JSON(todos)
}

func GetTODOById(c *fiber.Ctx) {
	id := c.Params("id")

	db := database.DB
	var todo Todo
	db.Find(&todo, id)
	if todo.Title == "" {
		c.Status(500).Send("No TODO Found with ID")
		return
	}
	c.JSON(todo)
}

func SaveTODO(c *fiber.Ctx) {

	db := database.DB
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&todo)
	c.JSON(todo)
}

func DeleteTODOById(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var todo Todo
	db.First(&todo, id)
	if todo.Title == "" {
		c.Status(500).Send("No TODO Found with ID")
		return
	}
	db.Delete(&todo)
	c.Send("TODO Successfully deleted")
}
