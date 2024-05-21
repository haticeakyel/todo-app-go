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

func(a *Api) HandleGetTodo(c *fiber.Ctx) error {
	ID := c.Params("id")
	todo, err := a.Service.GetTodo(ID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	c.JSON(todo)
	c.Status(fiber.StatusOK)
	return nil
}

func (a *Api) HandleGetTodos (c *fiber.Ctx) error{
	todos, err := a.Service.GetTodos()
	switch err{
	case nil:
		c.JSON(todos)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}

func (a *Api) HandleUpdateTodo(c *fiber.Ctx) error {
	ID := c.Params("id")

	todo := model.TodoDTO{}
	err := c.BodyParser(&todo)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	editTodo, err := a.Service.UpdateTodo(todo, ID)

	switch err {
	case nil:
		c.JSON(editTodo)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}
