run-gb-tag-create-consumer:
	ENV=local go run cmd/gb-tag-create-consumer/main.go

run-gb-tag-delete-consumer:
	ENV=local go run cmd/gb-tag-delete-consumer/main.go

compose-up:
	COMPOSE_PROJECT_NAME=garnbarn docker-compose up -d

compose-down:
	COMPOSE_PROJECT_NAME=garnbarn docker-compose down