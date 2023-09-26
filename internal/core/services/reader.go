package services

import (
	"time"

	"github.com/thalesb16/kafka-template/internal/core/domain"
	"github.com/thalesb16/kafka-template/internal/core/ports"
)

type readerService struct {
	queue ports.Queue
}

func NewReader(queue ports.Queue) *readerService {
	return &readerService{
		queue: queue,
	}
}

func (srv *readerService) Read(t time.Duration) (domain.Message, error) {
	return srv.queue.Read(t)
}
