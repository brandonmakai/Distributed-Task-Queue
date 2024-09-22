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

//var ctx = context.Background()

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/tasks", services.GetTask)
	server.HandleFunc("POST /task", services.PostTask)

	log.Fatal(http.ListenAndServe(":8081", server))
	/*




	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println(pong)
	*/

}