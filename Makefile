.PHONY: build
build:
	go build -o flare main.go

.PHONY: up
up:
	docker compose -f ./docker/docker-compose.yml up -d

.PHONY: down
down:
	docker compose -f ./docker/docker-compose.yml down