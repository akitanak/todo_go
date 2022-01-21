package entities

import (
	"time"

	"github.com/google/uuid"
)

// Todo Entity
type Todo struct {
	id          uuid.UUID
	description string
	dueDate     time.Time
	isFinished  bool
}

// Create new Todo. It assigns id with random UUID.
func NewTodo(description string) (*Todo, error) {
	if err := ValidateDescription(description); err != nil {
		return nil, err
	}

	todo := Todo{
		id:          uuid.New(),
		description: description,
		isFinished:  false,
	}

	return &todo, nil
}

// Id returns id.
func (t Todo) Id() uuid.UUID {
	return t.id
}

// Description returns description.
func (t Todo) Description() string {
	return t.description
}

// SetDescription updates description.
// `error` is returned if the input was invalid.
func (t *Todo) SetDescrption(description string) error {
	if err := ValidateDescription(description); err != nil {
		return err
	}
	t.description = description
	return nil
}

// DueDate returns dueDate.
func (t Todo) DueDate() time.Time {
	return t.dueDate
}

// IsFinished returns isFinished
func (t Todo) IsFinished() bool {
	return t.isFinished
}
