package repositories

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/akitanak/todo_go/entities"
)

type TodoRepository struct {
	file     *os.File
	todoList []entities.Todo
}

// Initialize TodoRepository.
func InitTodoRepository(path string) (*TodoRepository, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("InitTodoRepository was failed. %w", err)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open Todo file. %w", err)
	}
	defer file.Close()

	todoList := make([]entities.Todo, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		todo, err := deserialize(line)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize TodoRepository. %w", err)
		}
		todoList = append(todoList, *todo)
	}

	return &TodoRepository{file: file, todoList: todoList}, nil
}

func (repo *TodoRepository) Add(todo entities.Todo) {
	repo.todoList = append(repo.todoList, todo)
}

func (repo *TodoRepository) List(excludeFinished bool) []entities.Todo {
	list := make([]entities.Todo, 0)
	for _, todo := range repo.todoList {
		if !excludeFinished || !todo.IsFinished() {
			list = append(list, todo)
		}
	}

	return list
}
