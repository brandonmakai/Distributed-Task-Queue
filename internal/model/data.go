package model

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Task struct {
	ID string `json:"id"`
	Type string `json:"type"`
	Payload map[string]interface{} `json:"payload"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"`
}

type Producer struct {
	Client redis.Client
	Ctx context.Context
}