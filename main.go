package main

import (
	"log"
	"net/http"

	"taskapi/internal/task"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /task", task.CreateHandler)
	mux.HandleFunc("GET /task/", task.StatusHandler)
	mux.HandleFunc("DELETE /task/", task.DeleteHandler)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
