package main

import (
	"log"
	"net/http"

	"github.com/brandonmakai/task-queue/internal/gateway/api"
)

func main() {
	gateway := http.NewServeMux()

	gateway.HandleFunc("GET /api/tests", api.GetAllTask)
	gateway.HandleFunc("/api/test",api.PostTask)
	log.Fatal(http.ListenAndServe(":8080", gateway))
}