package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"io"

	_"github.com/brandonmakai/task-queue/internal/workers"
	"github.com/brandonmakai/task-queue/internal/model"

)

func main() {
	client := &http.Client{}

	data := model.Task{ID: "1", Message: "Hello World!", User: "Heavenly"}

	fmt.Printf("Data: %v\n", data)
	jsonData, _ := json.Marshal(data)
	fmt.Printf("Json: %v\n", jsonData)
	reader := bytes.NewReader(jsonData)

	resp, err := client.Post("http://localhost:8080/api/test", "application/json", reader)
	if err != nil {
		fmt.Printf("Failed to post data to backend at: %v", err)
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	
	var dataObj model.Task
	err = json.Unmarshal(body, &dataObj)
	if err != nil {
		log.Fatalf("Failed to unmarshal data back into hashmap: %v", err)
	}

	log.Println(dataObj.Message)
}