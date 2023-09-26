package main

import (
	"log"

	"github.com/thalesb16/kafka-template/internal/core/services"
	"github.com/thalesb16/kafka-template/pkg/kafka"
)

func main() {
	c, err := kafka.NewConsumer()
	if err != nil {
		panic(err)
	}

	r := services.NewReader(c)
	for {
		msg, err := r.Read(-1)
		if err != nil {
			panic(err)
		}

		log.Println(string(msg))
	}
}
