package main

import (
	"log"
	"net/http"

	"github.com/brandonmakai/task-queue/internal/gateway/api"
)

func main() {
	gateway := http.NewServeMux()

	gateway.HandleFunc("GET /api/tasks/{id}", api.GetTaskByID)
	gateway.HandleFunc("/api/task", api.PostTask)
	gateway.HandleFunc("/api/enqueue", api.EnqueueTask)
	gateway.HandleFunc("/api/pop", api.PopTask)
	log.Fatal(http.ListenAndServe(":8080", gateway))
}