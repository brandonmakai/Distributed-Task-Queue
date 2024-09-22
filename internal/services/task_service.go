package services

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gotten Task!"))
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	task := make(map[string]interface{})
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB: 0,
		Password: "",
	})

	body, err := io.ReadAll(r.Body); if err != nil {
		log.Fatalf("Failed to read data: %v", err)
	}

	err = json.Unmarshal(body, &task); if err != nil {
		log.Fatalf("Failed to unmarshal data: %v", err)
	}

	err = rdb.HSet(ctx, "task:1", task).Err(); if err != nil {
		log.Fatalf("Failed to convert Task to Redis Hash: %v", err)
		panic(err)
	}

	w.Write([]byte("Success!"))
}