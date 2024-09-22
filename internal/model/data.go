package model 

import (
	"github.com/go-redis/redis/v8"
	"context"
)

type Task struct {
	ID string `json:"ID"`
	Message string `json:"Message"`
	UserID string `json:"UserID"`
}

type User struct {
	ID string
	Name string
}

type Producer struct {
	Client redis.Client
	Ctx context.Context
}