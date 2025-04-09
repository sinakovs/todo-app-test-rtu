package main

import (
	"net/http"
)

func routes() http.Handler {

	api := http.NewServeMux()

	api.HandleFunc("GET /todos", getTodos)
	api.HandleFunc("POST /todos", addTodo)
	api.HandleFunc("GET /todos/", getTodo)
	api.HandleFunc("DELETE /todos/", deleteTodo)
	api.HandleFunc("PATCH /todos/", toggleTodoStatus)
	api.HandleFunc("GET /check", getTodosFromFile)
	api.HandleFunc("POST /save", saveDataToFile)
	api.HandleFunc("POST /parallel", getTodosAndSave)
	api.HandleFunc("GET /parallel/file", getTodosFromFileParallel)

	return api
}
