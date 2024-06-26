package connect

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func Client() (client *redis.Client, err error) {
	ctx := context.TODO()
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ping, err := client.Ping(ctx).Result()
	fmt.Println(ping, err)
	return
}
