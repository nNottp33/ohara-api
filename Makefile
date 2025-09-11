SHELL := /bin/bash
CONTAINER ?= app
MAIN_FILE=./cmd/server.go
DOCS_DIR=./docs

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

swag:
	@command -v swag >/dev/null 2>&1 || { \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	}
	swag init -g $(MAIN_FILE) -o $(DOCS_DIR)