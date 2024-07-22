package main

import "fmt"

type todo struct {
	id      uint
	title   string
	content string
}

var todos []todo

func main() {
	addTodo(1, "Homework", "Math exercises")
	addTodo(2, "Chores", "Clean room")
	getTodos()
	deleteTodoByID(1)
	getTodos()
}

func addTodo(id uint, title string, content string) {
	t := todo{id: id, title: title, content: content}
	todos = append(todos, t)
}

func getTodos() {
	for _, t := range todos {
		fmt.Println(t.id, t.title, t.content)
	}
}

func deleteTodoByID(i uint) {
	for index, t := range todos {
		if t.id == i {
			todos = append(todos[:index], todos[index+1:]...)
			fmt.Println("Deleted:", t.title)
			return
		}
	}
	fmt.Println("Todo not found")
}
