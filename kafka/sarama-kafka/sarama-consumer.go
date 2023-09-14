package sarama_kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"go-kafka-redis/model"
	"sync"
)

type Consumer struct {
	consumer sarama.Consumer
	messages chan *model.Message
	done     chan struct{}
}

func (c *Consumer) Messages() <-chan *model.Message {
	return c.messages
}

func (c *Consumer) Close() error {
	close(c.done)

	// Close the Sarama consumer
	if err := c.consumer.Close(); err != nil {
		return err
	}
	return nil
}

func NewSaramaConsumer(brokerList, topicName string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{brokerList}, config) // TODO: fix brokerList
	if err != nil {
		// TODO: handle error
		return nil, err
	}

	partitions, err := consumer.Partitions(topicName)
	if err != nil {
		// TODO: handle error
	}
	// Create a channel for consuming messages
	messages := make(chan *model.Message)

	// Create a channel to signal when the consumer is done
	done := make(chan struct{})

	// Create a wait group to track partitions
	var wg sync.WaitGroup

	// Start consuming messages for each partition in a separate goroutine
	for _, partition := range partitions {
		wg.Add(1)
		go func(partition int32) {
			defer wg.Done()
			partitionConsumer, err := consumer.ConsumePartition(topicName, partition, sarama.OffsetOldest)
			if err != nil {
				// Handle error
				return
			}
			defer partitionConsumer.Close()

			for {
				select {
				case err := <-partitionConsumer.Errors():
					// TODO: Handle err
					fmt.Println(err)
					// Example: log.Printf("Kafka consumer error on partition %d: %v", partition, err)
				case msg := <-partitionConsumer.Messages():
					// Convert the Sarama message to your generic ConsumerMessage type
					consumerMsg := &model.Message{
						Id:   string(msg.Key),
						Name: string(msg.Value),
					}
					messages <- consumerMsg
				case <-done:
					// Partition consumer is done, exit the loop
					return
				}
			}
		}(partition)
	}

	// Start a goroutine to wait for all partition consumers to finish
	go func() {
		wg.Wait()
		close(messages)
		close(done)
	}()

	return &Consumer{
		consumer: consumer,
		messages: messages,
		done:     done,
	}, nil
}
