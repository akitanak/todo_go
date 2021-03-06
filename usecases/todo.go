package usecases

import (
	"fmt"
	"time"

	"github.com/akitanak/todo_go/entities"
)

type CreateTodoOptions struct {
	DueDate time.Time
}

type CreateTodoResponse struct {
	Description string
	DueDate     time.Time
	IsFinished  bool
}

// create new todo.
func CreateTodo(description string, options *CreateTodoOptions) (*CreateTodoResponse, error) {
	todo, err := entities.NewTodo(description)
	if err != nil {
		err = fmt.Errorf("CreateTodo: Creating Todo was failed. %w", err)
		return nil, err
	}

	if options != nil {
		todo.SetDueDate(options.DueDate)
	}

	return &CreateTodoResponse{
		Description: todo.Description(),
		DueDate:     todo.DueDate(),
		IsFinished:  todo.IsFinished(),
	}, nil
}
