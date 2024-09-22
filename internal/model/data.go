package model 

import (
	"github.com/go-redis/redis/v8"
	"context"
)

type Task struct {
	ID string 
	Message string 
	User string 
}

type Producer struct {
	Client redis.Client
	Ctx context.Context
}