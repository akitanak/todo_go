package usecases

import (
	"strings"
	"testing"
	"time"
)

func TestCreateTodo(t *testing.T) {
	type input struct {
		description string
		options     *CreateTodoOptions
	}
	tests := map[string]struct {
		input   input
		want    *CreateTodoResponse
		wantErr string
	}{
		"normal case": {
			input: input{
				description: "make a todo app.",
				options: &CreateTodoOptions{
					DueDate: createUTCDate(2022, time.January, 21),
				},
			},
			want: &CreateTodoResponse{
				Description: "make a todo app.",
				DueDate:     createUTCDate(2022, time.January, 21),
				IsFinished:  false,
			},
			wantErr: "",
		},
		"without option": {
			input: input{
				description: "make a todo app.",
				options:     nil,
			},
			want: &CreateTodoResponse{
				Description: "make a todo app.",
				IsFinished:  false,
			},
			wantErr: "",
		},
		"too long description": {
			input: input{
				description: strings.Repeat("a", 65),
				options: &CreateTodoOptions{
					DueDate: createUTCDate(2022, time.January, 21),
				},
			},
			want:    nil,
			wantErr: "CreateTodo: Creating Todo was failed.",
		},
	}

	for name, test := range tests {
		response, err := CreateTodo(test.input.description, test.input.options)

		if test.want != nil {
			if err != nil {
				t.Errorf(`%v - CreateTodo(%v, %v) error happen. %v`, name, test.input.description, test.input.options, err)
			}
			if response.Description != test.want.Description {
				t.Errorf(`%v - CreateTodo(%v, %v) description unmatched. got: %v, want: %v`, name, test.input.description, test.input.options, response.Description, test.want.Description)
			}
			if response.DueDate != test.want.DueDate {
				t.Errorf(`%v - CreateTodo(%v, %v) DueDate unmatched. got: %v, want: %v`, name, test.input.description, test.input.options, response.DueDate, test.want.DueDate)
			}
		}
		if test.wantErr != "" {
			if err == nil {
				t.Errorf(`%v - CreateTodo(%v, %v) DueDate unmatched. got: %v, want: %v`, name, test.input.description, test.input.options, err, test.wantErr)
			}
			if !strings.HasPrefix(err.Error(), test.wantErr) {
				t.Errorf(`%v - CreateTodo(%v, %v) DueDate unmatched. got: %v, want: %v`, name, test.input.description, test.input.options, err, test.wantErr)
			}
		}

	}
}

func createUTCDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
