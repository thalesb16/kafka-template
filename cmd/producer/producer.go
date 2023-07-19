package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/thalesb16/kafka-template/internal/app"
)

func main() {
	p, err := app.NewProducer()
	if err != nil {
		panic(err)
	}

	values := []string{"message"}
	topic := "topic"
	for _, w := range values {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(w),
		}, nil)
	}

	p.Flush(15 * 1000)
}
