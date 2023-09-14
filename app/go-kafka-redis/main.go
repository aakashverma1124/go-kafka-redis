package main

import (
	"flag"
	CacheFactory "go-kafka-redis/cache/factory"
	"go-kafka-redis/constant"
	KafkaFactory "go-kafka-redis/kafka/factory"
)

var (
	brokerList   = flag.String("kafka.brokers", "localhost:29092", "The comma separated list of brokers in the Kafka cluster")
	topicList    = flag.String("kafka.topics", "e3f.markets", "REQUIRED: comma separated list of the topics to consume")
	cacheAddress = flag.String("cache.address", "localhost:6379", "This is cache address")
)

func main() {
	flag.Parse()

	cacheFactory := &CacheFactory.CacheFactory{}
	cacheObject := cacheFactory.GetCache(*cacheAddress, constant.REDIS_CACHE)
	if cacheObject == nil {
		//logger.L.Panic("Cache object couldn't be initialised", logger.String("Cache Address", *cacheAddress))
	}

	kafkaFactory := &KafkaFactory.KafkaFactory{}
	kafkaConsumer, err := kafkaFactory.GetKafkaConsumer(*brokerList, *topicList, constant.SARAMA_KAFKA)
	if err != nil {
		//logger.L.Panic("Kafka Consumer couldn't be initialised", logger.String("KafkaTopics", *topicList), logger.String("KafkaBrokers", *brokerList))
	}

	app := Service(kafkaConsumer, cacheObject)
	go app.StartProcessing()
	<-make(chan int)

}
