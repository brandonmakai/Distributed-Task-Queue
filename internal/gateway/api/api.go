package api

import (
	"net/http"
    "io"
	"log"
)

func GetTest(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	dataChan := make(chan []byte)

	go func() {
		resp, err := client.Get("http://localhost:8081/task")
		if err != nil {
			log.Printf("Failed to get tasks: %v", err)
			dataChan <- []byte("Error fetching tasks")
			return
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to get read task data: %v", err)
			dataChan <- []byte("Error reading task")
			return
		}

		dataChan <- data
	}()
	
	w.Write(<-dataChan) 
} 

func PostTest(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	dataChan := make(chan []byte)
	go func() {
		body := r.Body
		resp, err := client.Post("http://localhost:8081/task", "application/json", body)
		if err != nil {
			log.Printf("Failed to post data to task servce: %v", err)
			dataChan <- []byte("Failed to post data to task service.")
			return
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read data from task service: %v", err)
			dataChan <- []byte("Failed to read data from task service.")
			return
		}
		
		dataChan <- data
	}()

	w.Write(<-dataChan)
}