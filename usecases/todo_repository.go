package usecases

import (
	"github.com/akitanak/todo_go/entities"
)

type UserRepository interface {
	Add(todo entities.Todo) error
	Get(index int) (entities.Todo, error)
	List() ([]entities.Todo, error)
	Update(todo entities.Todo) error
}
