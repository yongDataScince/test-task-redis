package setupdb

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Setup struct{}

func NewSetup(ctx *context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdb.Ping(*ctx).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("PING REDIS: ", pong)

	return rdb, nil
}
