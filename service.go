package main

import (
	"strings"

	"github.com/google/uuid"
	model "github.com/haticeakyel/back-end/models"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) Service {
	return Service{
		Repository: repository,
	}
}

func GenerateUUID(length int) string {
	uuid := uuid.New().String()
	uuid = strings.ReplaceAll(uuid, "-", "")

	if length < 1 {
		return uuid
	}
	if length > len(uuid) {
		length = len(uuid)
	}

	return uuid[0:length]
}

func (s *Service) CreateTodo (todoDto model.TodoDTO)(*model.Todo, error){
	todoCreate := model.Todo{
		ID: GenerateUUID(8),
		Name: todoDto.Name,
		Description: todoDto.Description,
	}
	todoCreated, err := s.Repository.CreateTodo(todoCreate)
	if err!= nil{
		return nil, err
	}

	return todoCreated,nil
}

func (s *Service) GetTodo(ID string)(*model.Todo, error){
	gotTodo, err := s.Repository.GetTodo(ID)

	if err != nil{
		return nil, err
	}

	return gotTodo, nil
}

func (s *Service) GetTodos() ([]model.Todo, error){
	todoList, err := s.Repository.GetTodos()

	if err != nil{
		return nil, err
	}

	return todoList,nil
}

func (s *Service) UpdateTodo(todoDTO model.TodoDTO, ID string) (model.Todo, error) {

	updatedTodo, err := s.Repository.editTodo(todoDTO, ID)
	if err != nil {
		return model.Todo{}, nil
	}

	return updatedTodo, nil
}