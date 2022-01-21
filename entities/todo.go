package entities

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id          uuid.UUID
	Description string
	DueDate     time.Time
	IsFinished  bool
}

func NewTodo(description string) (*Todo, error) {
	if err := ValidateDescription(description); err != nil {
		return nil, err
	}

	todo := Todo{
		Id:          uuid.New(),
		Description: description,
		IsFinished:  false,
	}

	return &todo, nil
}
