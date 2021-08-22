package todo

import (
	"time"

	"github.com/fatopato/go-fiber-assistant/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title     string    `json:title`
	Completed bool      `json:completed`
	DueTime   time.Time `json:dueTime`
}

func GetAllTODOs(c *fiber.Ctx) error {
	db := database.DB
	var todos []Todo
	db.Find(&todos)
	c.JSON(todos)
	return nil
}

func GetTODOById(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB
	var todo Todo
	db.Find(&todo, id)
	if todo.Title == "" {
		c.Status(500).Send([]byte("No TODO Found with ID"))
		return nil
	}
	c.JSON(todo)
	return nil
}

func SaveTODO(c *fiber.Ctx) error {

	db := database.DB
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return nil
	}
	db.Create(&todo)
	c.JSON(todo)
	return nil
}

func UpdateTODO(c *fiber.Ctx) error {

	db := database.DB
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return err
	}
	db.Save(&todo)
	c.JSON(todo)
	return nil
}

func DeleteTODOById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var todo Todo
	db.First(&todo, id)
	if todo.Title == "" && todo.ID == 0 {
		c.Status(500).SendString("No TODO Found with ID")
		return nil
	}
	db.Delete(&todo)
	c.SendString("TODO Successfully deleted")
	return nil
}

func CheckOverTimeById(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB
	var todo Todo
	db.Find(&todo, id)
	if todo.Title == "" {
		c.Status(500).SendString("No TODO Found with ID")
		return nil
	}
	c.JSON(todo.isOverTime())
	return nil
}

func (todo *Todo) isOverTime() bool {
	return time.Now().After(todo.DueTime) && !todo.Completed && !todo.DueTime.IsZero()
}

func CompleteTODOById(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB
	var todo Todo
	db.Find(&todo, id)
	if todo.Title == "" {
		c.Status(500).SendString("No TODO Found with ID")
		return nil
	}
	todo.Completed = true
	db.Save(&todo)
	c.JSON(todo)
	return nil
}

func UndoTODOById(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB
	var todo Todo
	db.Find(&todo, id)
	if todo.Title == "" {
		c.Status(500).SendString("No TODO Found with ID")
		return nil
	}
	todo.Completed = false
	db.Save(&todo)
	c.JSON(todo)
	return nil
}
