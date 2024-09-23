package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"net/http"
	"time"

	"github.com/brandonmakai/task-queue/internal/model"
	"github.com/go-redis/redis/v8"
)
var ctx = context.Background()

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/task-service/tasks/")
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	key := fmt.Sprintf("task:%v", id)
	resp := rdb.HGetAll(ctx, key)
	
	task, err := resp.Result(); if err != nil {
		log.Fatalf("Failed to get redis hash: %v", err)
	}

	/*
	var task model.Task 

	err = resp.Scan(&task); if err != nil {
		log.Fatalf("Failed to convert redis hash to task: %v", err)
	}

	taskJson, err := json.Marshal(task); if err != nil {
		log.Fatalf("Failed to convert task into JSON: %v", err)
	}
	*/

	taskJson, err := json.Marshal(task); if err != nil {
		log.Fatalf("Failed to convert task into JSON: %v", err)
	}

	w.Write(taskJson)
}

func PostTask(w http.ResponseWriter, r *http.Request) {
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

func EnqueueTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	rdb := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
			Password: "",
			DB: 0,
	})

	data, err := io.ReadAll(r.Body); if err != nil {
		log.Fatalf("Failed to read http body params: %v", err)
	}

	err = json.Unmarshal(data, &task); if err != nil {
		log.Fatalf("Failed to marshal body params into Task: %v", err)
	}

	taskJson, err := json.Marshal(task); if err != nil {
		log.Fatalf("Failed to convert Task back into JSON: %v", err)
	}

	key := fmt.Sprintf("task:%v\n", task.ID)
	err = rdb.LPush(ctx, key, taskJson).Err(); if err != nil {
		log.Fatalf("Failed to enqueue Task JSON: %v into redis: %v", taskJson, err)
	}

	w.Write([]byte("Success!"))
}

func PopTask(w http.ResponseWriter, r *http.Request) {
	rdb := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
			Password: "",
			DB: 0,
	})

	resp := rdb.BLPop(ctx, time.Second * 10)

	body, err := resp.Result()
	fmt.Printf("BODY: %v\n", body)
	fmt.Printf("ERROR: %v\n", err)

	w.Write([]byte("Success!"))	
}
