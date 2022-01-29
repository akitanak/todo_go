package repositories

import (
	"testing"
	"time"

	"github.com/akitanak/todo_go/entities"
)

func TestSerialize(t *testing.T) {
	tests := map[string]struct {
		original entities.Todo
		want     string
	}{
		"standard case": {
			original: *createTodo(t, "build a todo app.", emptyTime, false),
			want:     "- [ ] build a todo app.",
		},
		"with DueDate": {
			original: *createTodo(t, "build a todo app.", time.Date(2022, time.January, 30, 0, 0, 0, 0, time.UTC), false),
			want:     "- [ ] build a todo app.\t2022-01-30",
		},
		"finished todo": {
			original: *createTodo(t, "build a todo app.", time.Date(2022, time.January, 30, 0, 0, 0, 0, time.UTC), true),
			want:     "- [x] build a todo app.\t2022-01-30",
		},
	}

	for name, test := range tests {
		serialized := serialize(test.original)

		if serialized != test.want {
			t.Errorf("%v - got: %v, want: %v", name, serialized, test.want)
		}
	}
}

func TestDeserialize(t *testing.T) {
	tests := map[string]struct {
		original string
		want     entities.Todo
		wantErr  string
	}{
		"standard case": {
			original: "- [ ] build a todo app.",
			want:     *createTodo(t, "build a todo app.", emptyTime, false),
			wantErr:  "",
		},
		"with DueDate": {
			original: "- [ ] build a todo app.\t2022-02-01",
			want:     *createTodo(t, "build a todo app.", time.Date(2022, time.February, 1, 0, 0, 0, 0, time.UTC), false),
			wantErr:  "",
		},
		"finished Todo": {
			original: "- [x] build a todo app.\t2022-02-01",
			want:     *createTodo(t, "build a todo app.", time.Date(2022, time.February, 1, 0, 0, 0, 0, time.UTC), true),
			wantErr:  "",
		},
	}

	for name, test := range tests {
		deserialized, err := deserialize(test.original)

		if *deserialized != test.want {
			t.Errorf("%v - got: %v, want: %v", name, *deserialized, test.want)
		}

		if test.wantErr != "" {
			err := err.Error()
			if err != test.wantErr {
				t.Errorf("%v - got: %v, want: %v", name, err, test.wantErr)
			}
		}
	}
}

var emptyTime time.Time

func createTodo(t *testing.T, description string, dueDate time.Time, isFinished bool) *entities.Todo {
	todo, err := entities.NewTodo(description)
	if err != nil {
		t.Fatal(err)
	}

	if !dueDate.IsZero() {
		todo.SetDueDate(dueDate)
	}

	if isFinished {
		todo.Finish()
	}

	return todo
}
