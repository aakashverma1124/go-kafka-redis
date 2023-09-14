package main

import (
	"encoding/json"
	"fmt"
	"go-kafka-redis/cache"
	"go-kafka-redis/kafka"
)

type App struct {
	kafkaConsumer kafka.Consumer
	cache         cache.Cache
}

func Service(consumer kafka.Consumer, cache cache.Cache) *App {
	return &App{
		kafkaConsumer: consumer,
		cache:         cache,
	}
}

func (app *App) StartProcessing() {
	for message := range app.kafkaConsumer.Messages() {
		fmt.Println(message)
		stringMessage, _ := json.Marshal(message)
		err := app.cache.SetKey(message.Id, string(stringMessage), 0)
		if err != nil {
			fmt.Print(err)
		}

		value, err := app.cache.GetKey(message.Id)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Println(value)
		}

	}
}
