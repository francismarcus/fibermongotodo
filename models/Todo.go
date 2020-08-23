package models

import (
	"github.com/Kamva/mgm/v2"
)

// Todo model
type Todo struct {
	// Inherit __id, created_at, updated_at
	mgm.DefaultModel `bson:",inline"`

	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Done        bool   `json:"done" bson:"done"`
}

// CreateTodo is a wrapper that creates a new todo entry
func CreateTodo(title, description string) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
		Done:        false,
	}
}
