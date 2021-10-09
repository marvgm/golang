package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	"text/template"
)

type Task struct {
	Name string
	Done bool
}

func main()  {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/tasks", TasksHandler)
	http.ListenAndServe(":8887", nil)
}

func HelloHandler(writer http.ResponseWriter, request *http.Request)  {
	id := uuid.NewV4().String()
	fmt.Fprintf(writer, "Helllo GO. id %v", id)
}

func TasksHandler(writer http.ResponseWriter, request *http.Request)  {

	tmpl := template.Must(template.ParseFiles("server/tasks.html"))

	tasks := []Task{
		{"Task 1", false},
		{"Task 2", true},
	}

	tmpl.Execute(writer, tasks)

	fmt.Fprint(writer, "Helllo GO")
}