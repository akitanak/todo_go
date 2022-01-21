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
		initDesc := "build todo app."
		todo, err := NewTodo(initDesc)
		if err != nil {
			t.Errorf(`%v - SetDescription(%v) was failed in initial Todo generation. init_desc: %v`, name, test.input, initDesc)
		}

		err = todo.SetDescription(test.input)
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
		todo, err := NewTodo(desc)
		if err != nil {
			t.Errorf(`%v - SetDueDate(%v) was failed in initial Todo generation. init_desc: %v`, name, test.dueDate, desc)
		}

		err = todo.SetDueDate(test.dueDate)
		if err == nil {
			if todo.DueDate() != test.dueDate {
				t.Errorf(`%v - SetDueDate(%v) got: %v, want: %[2]v`, name, test.dueDate, todo.DueDate())
			}
		}
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
