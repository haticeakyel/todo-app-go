package main

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/haticeakyel/back-end/models"
)

type Api struct {
	Service *Service
}

func NewApi(service *Service) *Api {
	return &Api{
		Service: service,
	}
}

func (a *Api) HandleAddTodo(c *fiber.Ctx) error {
	todoDTO := model.TodoDTO{}
	err := c.BodyParser(&todoDTO)
	if err != nil{
		c.Status(fiber.StatusBadRequest)
		return err
	}

	todoCreate, err := a.Service.CreateTodo(todoDTO)
	switch err{
	case nil:
		c.JSON(todoCreate)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{"error": err.Error()})
	}
	return nil
}