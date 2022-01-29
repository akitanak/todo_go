package repositories

import (
	"strings"
	"testing"
)

func TestInitTodoRepository(t *testing.T) {
	tests := map[string]struct {
		path     string
		wantPath string
		wantErr  string
	}{
		"normal case": {
			path:     "../../testdata/repositories/file/todo.md",
			wantPath: "/testdata/repositories/file/todo.md",
			wantErr:  "",
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
		}
		if test.wantErr != "" {
			//
		}
	}
}
