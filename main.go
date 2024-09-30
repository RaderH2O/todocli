package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"raderh2o/todocli/fileoperations"
	"raderh2o/todocli/todo"
	"strconv"
)

// Error for quitting the app
type QuitApp int

func (QuitApp) Error() string {
	return "Exited"
}

// Error for giving the wrong index/no index for some operations
type InvalidIndex int

func (val InvalidIndex) Error() string {
	return fmt.Sprintf("Index %d is invalid for that operation", val)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func processInput(input string) (byte, int, error) {
	var err error = nil
	index := 0
	if len(input) > 1 {
		index, err = strconv.Atoi(input[1:])
	}
	// Converting 1-indexed to 0-indexed
	index--

	return input[0], index, err
}

func executeOperation(scanner *bufio.Scanner, operation byte, index int, todos todo.Todos) (todo.Todos, error) {
	switch operation {
	case 'a':
		// Add a todo element
		fmt.Println("Enter the content of the Todo")
		if !scanner.Scan() {
			panic(scanner.Err())
		}

		content := scanner.Text()
		err := scanner.Err()
		check(err)
		todos = append(todos, todo.Todo{Content: content, Done: false})
	case 't':
		// Toggle a todo
		if index < 0 || index > len(todos)-1 {
			return todos, InvalidIndex(index + 1)
		}
		todos[index].Done = !todos[index].Done
	case 'd':
		// Remove a todo
		if index < 0 || index > len(todos)-1 {
			return todos, InvalidIndex(index + 1)
		}
		todos = append(todos[:index], todos[index+1:]...)
	case 'q':
		// Quit from the application
		return todos, QuitApp(0)
	}

	return todos, nil
}

func printTodos(todos todo.Todos) {
	for i, todo := range todos {
		fmt.Printf("%d %v\n", i+1, todo)
	}
}

func main() {

	if _, err := os.Stat("todo.txt"); errors.Is(err, os.ErrNotExist) {
		os.WriteFile("todo.txt", []byte{}, 0666)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var input string

	// Initial loading
	loaded, err := fileoperations.ReadFromFile("todo.txt")
	check(err)

	// Convert the file content to todo.Todos
	todos := todo.GetTodos(loaded)

	// Main program loop
	running := true
	for running {
		printTodos(todos)
		fmt.Println("(a)dd/(t)oggle/(d)elete/(q)uit")
		scanner.Scan()
		input = scanner.Text()

		operation, index, err := processInput(input)
		if err != nil {
			fmt.Println("Please enter a valid operation")
			continue
		}

		todos, err = executeOperation(scanner, operation, index, todos)

		switch t := err.(type) {
		case QuitApp:
			running = false
		case InvalidIndex:
			fmt.Println(t)
		}

		if running {
			fileoperations.WriteToFile("todo.txt", todos)
		}
	}
}
