package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	repository, err := NewRepository()
	if err != nil {
		log.Fatal(err)
	}
	defer repository.client.Disconnect(context.Background())
	service := NewService(repository)
	api := NewApi(&service)
	app := SetupApp(api)
	app.Listen(":3001")
}

func SetupApp(api *Api) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Post("/addTodo",api.HandleAddTodo)
	app.Get("/todos",api.HandleGetTodos)
	app.Get("/todos/:id",api.HandleGetTodo)
	app.Put("todos/:id",api.HandleUpdateTodo)
	
	return app
}
