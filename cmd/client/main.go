package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"io"
	"log"
	_"time"
	"fmt"

	_"github.com/brandonmakai/task-queue/internal/workers"
	"github.com/brandonmakai/task-queue/internal/model"
)

func main() {
	client := &http.Client{}

	payload := map[string]interface{}{"student_id": "1", "student_name": "brandon"}

	task := model.NewTask("1", "email", payload)

	jsonData, _ := json.Marshal(task)
	reader := bytes.NewReader(jsonData)

	resp, err := client.Post("http://localhost:8080/api/enqueue", "application/json", reader)
	if err != nil {
		log.Fatalf("Failed to post data to backend at: %v", err)
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	if string(body) == "Success!" {
		id := 1
		url := fmt.Sprintf("http://localhost:8080/api/tasks/%v", id)

		resp, err := client.Get(url); if err != nil {
			log.Fatalf("Failed to get task from redis: %v", err)
		}

		jsonResponse, err := io.ReadAll(resp.Body); if err != nil {
			log.Fatalf("Failed to read response data: %v", err)
		}
		
		var task model.Task 

		err = json.Unmarshal(jsonResponse, &task); if err != nil {
			log.Fatalf("Failed to marshal data back into Task: %v", err)
		}
	}
}