package todo

import (
	"fmt"
	"strings"
)

type Todo struct {
	Content string
	Done    bool
}

type Todos []Todo

func (todo Todo) String() string {
	checkmark := " "
	if todo.Done {
		checkmark = "x"
	}

	return fmt.Sprintf("- [%v] %v", checkmark, todo.Content)
}

func (todos Todos) String() string {
	result := ""
	for _, todo := range todos {
		result += todo.String()
		result += "\n"
	}
	return result
}

func GetTodos(input string) Todos {
	lines := strings.Split(input, "\n")
	todos := Todos{}

	for _, line := range lines {
		parsed := ""
		var currentTodo Todo
		for i, c := range line {
			parsed += string(c)
			if parsed == "- [ ] " {
				currentTodo.Done = false
				currentTodo.Content = strings.Replace(line[i+1:], "\r", "", -1)
				break
			} else if parsed == "- [x] " {
				currentTodo.Done = true
				currentTodo.Content = strings.Replace(line[i+1:], "\r", "", -1)
				break
			}
		}
		if currentTodo.Content != "" {
			todos = append(todos, currentTodo)
		}
	}
	return todos
}

// func AddTodo(todos Todos, content string, status bool) Todos {
// 	return append(todos, Todo{Content: content, Done: status})
// }
