env-up:
	docker compose up -d

run-consumer:
	go run cmd/consumer/consumer.go
