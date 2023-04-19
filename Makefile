run-gb-tag-create-consumer:
	ENV=local go run cmd/gb-tag-create-consumer/main.go

compose-up:
	COMPOSE_PROJECT_NAME=garnbarn docker-compose up -d

compose-down:
	COMPOSE_PROJECT_NAME=garnbarn docker-compose down