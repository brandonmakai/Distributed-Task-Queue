package api

import (
	"net/http"
    "io"
	"log"
	"strings"
	"fmt"
)

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	dataChan := make(chan []byte)

	id := strings.TrimPrefix(r.URL.Path, "/api/tasks/")
	log.Printf("ID: %v", id)
	url := fmt.Sprintf("http://localhost:8081/task-service/tasks/%v", id)

	go func() {
		resp, err := client.Get(url)
		if err != nil {
			dataChan <- []byte("Error fetching tasks")
			log.Fatalf("Failed to get tasks: %v", err)
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			dataChan <- []byte("Error reading task")
			log.Fatalf("Failed to get read task data: %v", err)
		}

		dataChan <- data
	}()
	
	w.Write(<-dataChan) 
} 

func PostTask(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	dataChan := make(chan []byte)

	go func() {
		body := r.Body
		resp, err := client.Post("http://localhost:8081/task-service/task", "application/json", body)
		if err != nil {
			dataChan <- []byte("Failed to post data to task service.")
			log.Fatalf("Failed to post data to task servce: %v", err)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			dataChan <- []byte("Failed to read data from task service.")
			log.Fatalf("Failed to read data from task service: %v", err)
		}
		
		dataChan <- data
	}()

	w.Write(<-dataChan)
}

func EnqueueTask(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	dataChan := make(chan []byte)

	go func() {
		resp, err := client.Post("http://localhost:8081/task-service/enqueue", "application/json", r.Body); if err != nil {
			dataChan <- []byte("Failed to post (enqueue) data to task service.")
			log.Fatalf("Failed to post (enqueue) data to task service: %v", err)
		}

		data, err := io.ReadAll(resp.Body); if err != nil {
			dataChan <- []byte("Failed to read data from task service.")
			log.Fatalf("Failed to read data from task service: %v", err)
		}
		
		dataChan <- data
	}()

	w.Write(<-dataChan)
}

func PopTask(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}

	dataChan := make(chan []byte)
	go func() {
		resp, err := client.Get("localhost:8081/task-service/pop"); if err != nil {
			dataChan <- []byte("Failed to pop task from task service.")
			log.Fatalf("Failed to pop task from task service: %v", err)
		}

		body, err := io.ReadAll(resp.Body); if err != nil {
			dataChan <- []byte("Failed to read data from task service.")
			log.Fatalf("Failed to read task service (pop) response: %v", err)
		}

		dataChan <- body
	}()

	w.Write(<-dataChan)
}