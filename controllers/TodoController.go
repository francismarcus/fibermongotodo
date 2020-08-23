package controllers

import (
	"github.com/Kamva/mgm/v2"
	"github.com/francismarcus/fibermongotodo/models"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllTodos - GET /api/todos
func GetAllTodos(ctx *fiber.Ctx) {
	//accessing the todos collection
	collection := mgm.Coll(&models.Todo{})

	//Initialize an empty array to store our todos
	todos := []models.Todo{}

	/*
			collection.SimpleFind()
		 	The function takes two parameters:
		 	the first parameter is the memory address of the variable in which the result should be stored
			the second is a filter. If the filter is empty, it will return all entries.
	*/

	err := collection.SimpleFind(&todos, bson.D{})
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		// Important to return else the controller will continue
		return
	}
	ctx.JSON(fiber.Map{
		"ok":    true,
		"todos": todos,
	})

}

// GetTodoByID - GET /api/todos/:id
func GetTodoByID(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found",
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

// CreateTodo - POST /api/todos
func CreateTodo(ctx *fiber.Ctx) {
	// create a new struct that contains the parameters we need to extract from the request’s body.
	params := new(struct {
		Title       string
		Description string
	})

	// We can then use Fiber’s ctx.BodyParser() method to parse the request’s body and bind it to our params variable:
	ctx.BodyParser(&params)

	// We can then check if both parameters were provided, and if any of them is missing, we can return an error with an HTTP status code of 400 (Bad Request):
	if len(params.Title) == 0 || len(params.Description) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Title or description not specified.",
		})
		return
	}

	todo := models.CreateTodo(params.Title, params.Description)
	err := mgm.Coll(todo).Create(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

// ToggleTodoStatus - PATCH /api/todos/:id
func ToggleTodoStatus(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return
	}

	todo.Done = !todo.Done

	err = collection.Update(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

// DeleteTodo - DELETE /api/todos/:id
func DeleteTodo(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return
	}

	err = collection.Delete(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}
