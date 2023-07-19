package app

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewConsumer() (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:19092",
		"group.id":          "group",
	})
	if err != nil {
		return nil, err
	}

	c.SubscribeTopics([]string{"topic"}, nil)

	return c, nil
}
