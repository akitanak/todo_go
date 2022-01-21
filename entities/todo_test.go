package entities

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewTodo(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    *Todo
		wantErr error
	}{
		"create new Todo.": {
			input:   "build todo app with golang.",
			want:    &Todo{description: "build todo app with golang."},
			wantErr: nil,
		},
		"too long description": {
			input:   strings.Repeat("a", 65),
			want:    nil,
			wantErr: errors.New("description is too long. max: 64, actual: 65"),
		},
	}

	for name, test := range tests {
		todo, err := NewTodo(test.input)

		if todo != nil && test.want != nil {
			assertTodo(t, name, test.input, *todo, *test.want)
		}

		if err != test.wantErr && err.Error() != test.wantErr.Error() {
			t.Errorf(`%v - NewTodo(%v) returns error. got: \"%v\", want: \"%v\"`, name, test.input, err.Error(), test.wantErr.Error())
		}
	}
}

func TestSetDescription(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    string
		wantErr string
	}{
		"valid description": {
			input:   "Build Todo App with Golang.",
			want:    "Build Todo App with Golang.",
			wantErr: "",
		},
		"too long description": {
			input:   strings.Repeat("a", 65),
			want:    "",
			wantErr: "description is too long. max: 64, actual: 65",
		},
	}

	for name, test := range tests {
		desc := "build todo app."
		todo := createInitialTodo(t, name, desc)

		err := todo.SetDescription(test.input)
		if test.want != "" {
			if todo.Description() != test.want {
				t.Errorf(`%v - SetDescription(%v) got: %v, want: %v`, name, test.input, todo.Description(), test.want)
			}
		}

		if test.wantErr != "" {
			if err.Error() != test.wantErr {
				t.Errorf(`%v - SetDescription(%v) got: %v, want: %v`, name, test.input, err.Error(), test.wantErr)
			}
		}
	}
}

func TestSetDueDate(t *testing.T) {
	tests := map[string]struct {
		dueDate time.Time
	}{
		"valid dueDate": {
			dueDate: time.Now().UTC(),
		},
	}

	for name, test := range tests {
		desc := "make todo app."
		todo := createInitialTodo(t, name, desc)

		err := todo.SetDueDate(test.dueDate)
		if err == nil {
			if todo.DueDate() != test.dueDate {
				t.Errorf(`%v - SetDueDate(%v) got: %v, want: %[2]v`, name, test.dueDate, todo.DueDate())
			}
		}
	}
}

func TestFinish(t *testing.T) {
	name := "finish Todo"
	desc := "make todo app."
	todo := createInitialTodo(t, name, desc)

	todo.Finish()

	if !todo.IsFinished() {
		t.Errorf(`%v - Finish() failed. got: %v, want: %v`, name, true, todo.IsFinished())
	}

}

func assertTodo(t *testing.T, name string, input string, got, want Todo) {
	var zeroUuid uuid.UUID
	if got.Id() == zeroUuid {
		t.Errorf(`%v - NewTodo(%v).Id() got: \"%v\", want: \"%v\"`, name, input, got.Id(), "non zero UUID")
	}

	if got.Description() != want.Description() {
		t.Errorf(`%v - NewTodo(%v).Description() got: \"%v\", want: \"%v\"`, name, input, got.Description(), want.Description())
	}

	if got.IsFinished() != false {
		t.Errorf(`%v - NewTodo(%v).IsFinished() got: \"%v\", want: \"%v\"`, name, input, got.IsFinished(), false)
	}
}

func createInitialTodo(t *testing.T, name string, description string) *Todo {
	todo, err := NewTodo(description)
	if err != nil {
		t.Errorf(`%v - failed in initial Todo creation. description: %v`, name, description)
	}
	return todo
}
