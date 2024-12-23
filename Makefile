# Simple Makefile for a Go project

# Build the application

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# confirm: display confirmation prompt
.PHONY: confirm
confirm:
	@echo -n "Are you sure? [y/N] " && read ans && [ $${ans:-N} = y ]


## build: Build the application
build:
	@echo "Building..."
	
	
	@go build -o main cmd/api/main.go

## run: Run the application
run:
	@go run cmd/api/main.go

## docker-run: Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

## docker-down: Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi


## clean: Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

## watch: Live Reload
watch:
	nodemon -e go --exec "make run" --signal SIGTERM

## swagger: Generate swagger docs
.PHONY: swagger
swagger:
    @echo "Generating Swagger documentation..."
    swag init -g cmd/api/main.go


## migrations/create: create a database migration
.PHONY: migrations/create
migrations/create: confirm
	@echo "creating migration files for ${name}..."
	migrate create -seq -ext .sql -dir ./migrations ${name}

## migrations/version: display current database migration version
.PHONY: migrations/version
migrations/version:
	@echo -n "database migration version: "
	@migrate -path ./migrations -database ${DB_URI} version

## migrations/up: apply all up database migrations
.PHONY: migrations/up
migrations/up: 
	@echo "running up migrations..."
	@migrate -path ./migrations -database ${DB_URI} up

## migrations/reset: rollback all database migrations
.PHONY: migrations/reset
migrations/reset: confirm
	@echo "running down migrations..."
	@migrate -path ./migrations -database ${DB_URI} down

## migrations/goto version=$1: move to a specified database migration version
.PHONY: migrations/goto
migrations/goto: confirm
	@echo "migrating database to version ${version}..."
	@migrate -path ./migrations -database ${DB_URI} goto ${version}

## migrations/force version=$1: force a migration to the given version
.PHONY: migrations/force
migrations/force: confirm
	@echo "migrating database to version ${version}..."
	@migrate -path ./migrations -database ${DB_URI} force ${version}
