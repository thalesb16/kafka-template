env-up:
	docker compose up -d

env-down:
	docker-compose down

run-consumer:
	go run cmd/reader/reader.go

run-producer:
	go run cmd/writer/writer.go
