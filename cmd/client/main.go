package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"
	"io"
	"context"
	"log"

	_"github.com/brandonmakai/task-queue/internal/workers"
	"github.com/brandonmakai/task-queue/internal/model"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	user := model.User{ID: "1", Name: "Heavenly"}

	client := &http.Client{}

	data := model.Task{ID: "1", Message: "Hello World!", UserID: user.ID}

	fmt.Printf("Data: %v\n", data)
	jsonData, _ := json.Marshal(data)
	fmt.Printf("Json: %v\n", jsonData)
	reader := bytes.NewReader(jsonData)

	resp, err := client.Post("http://localhost:8080/api/test", "application/json", reader)
	if err != nil {
		fmt.Printf("Failed to post data to backend at: %v", err)
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response String: %v", string(body))
	
	if string(body) == "Success!" {
		resp, err := rdb.HGet(ctx, "task:1", "Message").Result(); if err != nil {
			log.Printf("Failed to get hash message: %v", err)
		}
		log.Printf("Response Data: %v", resp)
	}
}