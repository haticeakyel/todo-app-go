package model

type Todo struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Done bool   `json:"done" bson:"done"`
}

type TodoDTO struct {
	Name string `json:"name" bson:"name"`
	Done bool   `json:"done" bson:"done"`
}
