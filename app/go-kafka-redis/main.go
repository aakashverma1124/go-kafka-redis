package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-kafka-redis/model"
)

func main() {
	fmt.Println("Hello, World")

	ctx := context.TODO()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// testing the redis connection
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	// setting value to key
	err = client.Set(ctx, "name", "Aakash", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	// getting value from key
	value, err := client.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)

	// example to test with composite object
	json, err := json.Marshal(model.Author{Name: "Elliot", Age: 25})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(ctx, "id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get(ctx, "id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

}
