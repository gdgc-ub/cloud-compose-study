# variables
include .env
export $(shell sed 's/=.*//' .env)
DB_URL = "postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

.PHONY: help
help:
	@echo "Choose a command:"
	@echo "  make run            	- Run the Go application"
	@echo "  make air            	- Run the Go application with live reloading"
	@echo "  make test           	- Run Go unit tests defined in ./tests"
	@echo "  make migrate-up     	- Apply all migrations"
	@echo "  make migrate-down   	- Rollback all migrations"
	@echo "  make create-migration name=<name> 	- Create a new migration file"

.PHONY: run
run:
	go run ./app/cmd/main.go

.PHONY: air
air:
	air -c .air.toml

.PHONY: compose-up
compose-up:
ifndef file
	docker compose up -d
else
	docker compose -f $(file) up -d
endif 

.PHONY: compose-down
compose-down:
ifndef file
	docker compose down
else
	docker compose -f $(file) down 
endif 

.PHONY: test
test:
	go test ./tests/...

.PHONY: migrate-up
migrate-up:
	migrate -path ./data/db/migrations -database ${DB_URL} up

.PHONY: migrate-down
migrate-down:
	migrate -path ./data/db/migrations -database ${DB_URL} down 

.PHONY: migrate-force
migrate-force:
ifndef version 
	$(error "Migration version not specified. Use 'make migrate-force version=<version>'")
endif 
	migrate -path ./data/db/migrations -database ${DB_URL} force $(version)

.PHONY: migrate-create
migrate-create:
ifndef name 
	$(error "Migration name not specified. Use 'make create-migration name=<name>'")
endif 
	migrate create -ext sql -dir ./migrations -seq $(name)