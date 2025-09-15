SHELL := /bin/bash
CONTAINER ?= app
MAIN_FILE=./cmd/server.go
DOCS_DIR=./docs

run:
	source .env && air -c .air.toml

build:
	go build -o ./tmp/main ./cmd/server.go

up:
	docker compose --env-file .env up --build -d

down:
	docker compose --env-file .env down

logs:
	docker compose logs -f $(CONTAINER)

prune:
	docker system prune -af --volumes

clean:
	rm -rf tmp/

swag:
	@command -v swag >/dev/null 2>&1 || { \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	}
	swag init -g $(MAIN_FILE) -o $(DOCS_DIR)