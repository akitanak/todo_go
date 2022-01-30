package repositories

import (
	"strings"
	"testing"
	"time"

	"github.com/akitanak/todo_go/entities"
)

func TestInitTodoRepository(t *testing.T) {
	tests := map[string]struct {
		path     string
		wantPath string
		wantList []entities.Todo
	}{
		"initial execution case": {
			path:     "../../testdata/repositories/file/todo.md",
			wantPath: "/testdata/repositories/file/todo.md",
			wantList: make([]entities.Todo, 0),
		},
		"todo file loading case": {
			path:     "../../testdata/repositories/file/standard_todo_file.md",
			wantPath: "/testdata/repositories/file/standard_todo_file.md",
			wantList: []entities.Todo{
				*createTodo(t, "make a todo app.", zeroValueTime, false),
				*createTodo(t, "pay car pool fee.", time.Date(2022, time.January, 30, 0, 0, 0, 0, time.UTC), true),
				*createTodo(t, "send books.", time.Date(2022, time.February, 1, 0, 0, 0, 0, time.UTC), false),
			},
		},
	}

	for name, test := range tests {
		repo, err := InitTodoRepository(test.path)

		if test.wantPath != "" {
			if err != nil {
				t.Errorf("%v - InitTodoRepository was failed unexpectedly. %w", name, err)
			}
			if !strings.HasSuffix(repo.file.Name(), test.wantPath) {
				t.Errorf("%v - file path was unmatched. got: %v, want: %v", name, repo.file.Name(), test.wantPath)
			}
			if len(repo.todoList) != len(test.wantList) {
				t.Errorf("%v - todoList len unmatched. got: %v, want: %v", name, len(repo.todoList), len(test.wantList))
			}
			for i, todo := range repo.todoList {
				if todo != test.wantList[i] {
					t.Errorf("%v - todo item is unmatched. got: %v, want: %v", name, todo, test.wantList[i])
				}
			}
		}
	}
}
