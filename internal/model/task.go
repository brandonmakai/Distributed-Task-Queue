package model

import (
	"time"
)

func NewTask(ID string, Type string, Payload map[string]interface{}) *Task {
	task := Task{
		ID: ID, 
		Type: Type, 
		Payload: Payload,
		Status: "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &task
}