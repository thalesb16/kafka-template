package ports

import (
	"github.com/thalesb16/kafka-template/pkg/kafka"
)

type Queue interface {
	kafka.Kafka
}
