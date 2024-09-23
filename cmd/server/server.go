package main

import (
	_"context"
	_"fmt"
	"log"

	_ "github.com/brandonmakai/task-queue/internal/model"
	"github.com/brandonmakai/task-queue/internal/services"
	_ "github.com/brandonmakai/task-queue/internal/services"

	"net/http"

	_"github.com/go-redis/redis/v8"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/task-service/tasks/{id}", services.GetTask)
	server.HandleFunc("/task-service/task", services.PostTask)
	server.HandleFunc("/task-service/enqueue", services.EnqueueTask)
	server.HandleFunc("/task-service/pop", services.PopTask)

	log.Fatal(http.ListenAndServe(":8081", server))
}