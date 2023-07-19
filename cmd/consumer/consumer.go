package main

import (
	"log"

	"github.com/thalesb16/kafka-template/internal/app"
)

func main() {
	c, err := app.NewConsumer()
	if err != nil {
		panic(err)
	}

	defer c.Close()

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			panic(err)
		}

		log.Println(string(msg.Value))
	}
}
