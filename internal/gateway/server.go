package main

import (
	"log"
	"net/http"

	"github.com/brandonmakai/task-queue/internal/gateway/api"
)

func main() {
	gateway := http.NewServeMux()

	gateway.HandleFunc("GET /api/tests", api.GetTest)
	gateway.HandleFunc("POST /api/test",api.PostTest )
	log.Fatal(http.ListenAndServe(":8080", gateway))
}