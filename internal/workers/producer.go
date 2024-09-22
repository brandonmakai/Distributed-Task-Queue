package workers

import (
	_"github.com/brandonmakai/task-queue/internal/model"
	_"github.com/go-redis/redis/v8"
	_"context"
	"net/http"

)


var Client http.Client

/*

func (p *model.Producer) Enqueue(t model.Task) {

}

func Enqueue(rdb redis.Client, ctx context.Context, task model.Task) {
	val := map[string]string{
		"user" : task.User, 
		"message" : task.Message,
	}
	rdb.RPush(ctx, task.ID, val)
}

func Dequeue()
*/