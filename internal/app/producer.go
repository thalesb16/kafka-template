package app

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

func NewProducer() (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:19092",
	})
	if err != nil {
		return nil, err
	}

	return p, nil
}
