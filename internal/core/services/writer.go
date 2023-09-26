package services

import "github.com/thalesb16/kafka-template/internal/core/ports"

type writerService struct {
	queue ports.Queue
}

func NewWriter(queue ports.Queue) *writerService {
	return &writerService{
		queue: queue,
	}
}

func (srv *writerService) Write(msg []string) error {
	return srv.queue.Write(msg)
}
