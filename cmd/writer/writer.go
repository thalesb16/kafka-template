package main

import (
	"github.com/thalesb16/kafka-template/internal/core/services"
	"github.com/thalesb16/kafka-template/pkg/kafka"
)

func main() {
	p, err := kafka.NewProducer()
	if err != nil {
		panic(err)
	}

	w := services.NewWriter(p)
	msg := []string{"example message"}

	w.Write(msg)
}
