package factory

import (
	"errors"
	"go-kafka-redis/constant"
	"go-kafka-redis/kafka"
	"go-kafka-redis/kafka/sarama-kafka"
)

type KafkaFactory struct{}

func (kafkaFactory *KafkaFactory) GetKafkaConsumer(brokerList, topicName, consumerType string) (kafka.Consumer, error) {
	if consumerType == constant.SARAMA_KAFKA {
		return sarama_kafka.NewSaramaConsumer(brokerList, topicName)
	}
	return nil, errors.New("couldn't get kafka consumer")
}
