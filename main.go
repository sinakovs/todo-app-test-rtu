package main

import (
	"net/http"
	"time"
)

type todo struct {
	TodosNumber int    `json;"id"`
	ID          string `json;"id"`
	Item        string `json;"item"`
	Completed   bool   `json;"completed"`
}

var todos, _ = getTodoDataFromFile()
var currentTodosInFile = len(*todos)

func main() {
	startWorkerPool(10)

	port := "8080"

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        routes(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	server.ListenAndServe()
}
