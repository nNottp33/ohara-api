SHELL := /bin/bash
CONTAINER ?= app

run:
	source .env && air -c .air.toml

build:
	go build -o ./tmp/main ./cmd/server.go

up:
	podman-compose --env-file .env up --build -d

down:
	podman-compose --env-file .env down

logs:
	podman-compose logs -f $(CONTAINER)

prune:
	podman system prune -af --volumes

clean:
	rm -rf tmp/