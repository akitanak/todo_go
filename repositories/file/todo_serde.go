package repositories

import (
	"fmt"
	"strings"
	"time"

	"github.com/akitanak/todo_go/entities"
)

// Serialize entities.Todo to Markdown check list format.
func serialize(todo entities.Todo) string {

	isFinished := serializeIsFinished(todo.IsFinished())
	if todo.DueDate().IsZero() {
		return fmt.Sprintf("- [%v] %v", isFinished, todo.Description())
	}

	dueDate := todo.DueDate().Format(dueDateLayout)
	return fmt.Sprintf("- [%v] %v\t%v", isFinished, todo.Description(), dueDate)
}

func serializeIsFinished(isFinished bool) string {
	if isFinished {
		return "x"
	}
	return " "
}

// Deserialize Markdown check list format to entities.Todo
func deserialize(row string) (*entities.Todo, error) {
	if !hasChecklist(row) {
		return nil, fmt.Errorf("invalid format. row must be task list styled. %v", row)
	}

	cols := strings.Split(row, "\t")

	description := extractDescription(cols[0])
	todo, err := entities.NewTodo(description)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize Todo. %w", err)
	}

	isFinished := isFinished(cols[0])
	if isFinished {
		todo.Finish()
	}

	if len(cols) > 1 {
		dueDate, err := parseDueDate(cols[1])
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize Todo. %w", err)
		}
		todo.SetDueDate(dueDate)
	}

	return todo, nil
}

func hasChecklist(row string) bool {
	return strings.HasPrefix(row, "- [ ]") || strings.HasPrefix(row, "- [x]")
}

func isFinished(row string) bool {
	taskStatus := string(row[3])
	switch taskStatus {
	case " ":
		return false
	case "x":
		return true
	default:
		panic(row)
	}
}

func extractDescription(row string) string {
	return row[6:len(row)]
}

var dueDateLayout = "2006-01-02"

func parseDueDate(col string) (time.Time, error) {
	dueDate, err := time.Parse(dueDateLayout, col)
	if err != nil {
		return dueDate, fmt.Errorf("invalid Due date format. DueDate must be YYYY-MM-dd. %w", err)
	}

	return dueDate, nil
}
