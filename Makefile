env-up:
	docker compose up -d

env-down:
	docker-compose down

run-consumer:
	go run cmd/consumer/consumer.go

run-producer:
	go run cmd/producer/producer.go
