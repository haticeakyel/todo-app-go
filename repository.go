package main

import (
	"context"
	"fmt"
	"log"
	"time"

	model "github.com/haticeakyel/back-end/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

func NewRepository() (*Repository, error) {
	uri := "mongodb+srv://haticeeakyel:test123@cluster0.6uwmyxf.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Repository{client}, nil
}

func (repository Repository) CreateTodo(todoData model.Todo) (*model.Todo, error) {
	collection := repository.client.Database("todo").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, todoData)
	if err != nil {
		return nil, fmt.Errorf("failed to insert todo: %w", err)
	}

	return &todoData, nil
}

func (repository Repository) GetTodos() ([]model.Todo, error) {
	collection := repository.client.Database("todo").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	todos := []model.Todo{}
	for cur.Next(ctx) {
		var todo model.Todo
		err := cur.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)

	}

	return todos, nil
}

func (repository Repository) GetTodo(ID string) (*model.Todo, error) {
	collection := repository.client.Database("todo").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	todo := &model.Todo{}
	filters := bson.M{"id": ID}
	err := collection.FindOne(ctx, filters).Decode(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (repository *Repository) editTodo(todoDTO model.TodoDTO, ID string) (model.Todo, error) {
	collection := repository.client.Database("todo").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateTodo := bson.M{
		"id":          ID,
		"name":        todoDTO.Name,
		"description": todoDTO.Description,
	}

	_, err := collection.ReplaceOne(ctx, bson.M{"id": ID}, updateTodo)

	if err != nil {
		return model.Todo{}, err
	}

	updatedTodo, err := repository.GetTodo(ID)

	if err != nil {
		return model.Todo{}, err
	}

	return *updatedTodo, nil
}

func (repository *Repository) DeleteTodo(ID string) error {
	collection := repository.client.Database("todo").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	deleteTodo := collection.FindOneAndDelete(ctx, bson.M{"id": ID})

	if deleteTodo != nil {
		return deleteTodo.Err()
	}

	return nil
}
