package model

type Todo struct {
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

}

type TodoDTO struct{
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}