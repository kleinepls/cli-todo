package main

import (
	"fmt"
	"os"
	"slices"
)

type Todo struct {
	id        int
	title     string
	completed bool
}

var todos = make([]Todo, 0)

func main() {
	if len(os.Args) == 1 {
		fail("no command provided.\navailable commands are: create, list, complete, delete.")
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case Create:
		handleCreate(args)
	default:
		fail("invalid command: '%s'.\navailable commands are: create, list, complete, delete.", command)
	}
}

func handleCreate(args []string) {
	if len(args) == 0 {
		fail("no task provided.")
	}
	if len(args) > 1 {
		fail("too many arguments provided.")
	}

	title := args[0]

	if slices.ContainsFunc(todos, func(todo Todo) bool {
		return todo.title == title
	}) {
		fail("task '%s' already exists.", title)
	}

	todos = append(todos, Todo{
		id:        len(todos),
		title:     title,
		completed: false,
	})

	var unfinishedTasks int
	for _, t := range todos {
		if !t.completed {
			unfinishedTasks++
		}
	}

	fmt.Printf("task '%s' created successfully.\n", title)
	if unfinishedTasks > 1 {
		fmt.Printf("you have %d unfinished tasks in total.\n", unfinishedTasks)
	}
}
