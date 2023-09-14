package kafka

import "go-kafka-redis/model"

type Consumer interface {
	Messages() <-chan *model.Message
	Close() error
}
