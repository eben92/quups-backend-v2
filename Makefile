# Simple Makefile for a Go project
include .env


DATABASE_URL="postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable"
GOOSE_MIGRATION_DIR= "./internal/database/migrations"
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=$(DATABASE_URL)


ec:
	@goose $(GOOSE_DRIVER) "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_DATABASE) sslmode=disable" status

# Build the application
all: build

build:
	@echo "Building..."
	
	@go build -o out/main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

schema:
	@echo "Generating schema..."
	@read -p "Enter schema name: " schema_name; \
	goose -dir $(GOOSE_MIGRATION_DIR) create $$schema_name sql  
	@echo "Generated!"

sqlc: 
	@echo "Generating..."
	@sqlc generate
	@echo "Generated!"

db-pull:
	@echo "Pulling..."
	@supabase db pull --db-url $(DATABASE_URL)
	@echo "Pulled!"	

db-pull-debug:
	@echo "Pulling..."
	@supabase db pull --debug
	@echo "Pulled!"		

db-push:
	@echo "Pushing..."
	@supabase db push
	@echo "Pushed!"

db-reset:
	@echo "Diffing..."
	@goose -dir "$(GOOSE_MIGRATION_DIR)" $(GOOSE_DRIVER) "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_DATABASE) sslmode=disable" reset
	@echo "Diffed!"

db-status:
	@echo "Status..."
	@goose -dir "$(GOOSE_MIGRATION_DIR)" $(GOOSE_DRIVER) "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_DATABASE) sslmode=disable" status
	@echo "DONE!"		

db-migrate-up:
	@echo "Migrating..."
	@goose -dir "$(GOOSE_MIGRATION_DIR)" $(GOOSE_DRIVER) "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_DATABASE) sslmode=disable" up
	@echo "Migrated!"

db-migrate-down:
	@echo "Migrating..."
	@goose -dir "$(GOOSE_MIGRATION_DIR)" $(GOOSE_DRIVER) "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_DATABASE) sslmode=disable" down
	@echo "Migrated!"

db-sync:
	@echo "syncing db..."
	@supabase migration up
	@echo "synced"

db-repair: 
	@echo "Repairing..."
	@read -p "Enter migration ID: " id; \
	read -p "Enter migration status (applied or reverted): " status; \
	supabase migration repair $$id --status $$status; \
	echo "Repaired!"

db-repair-raw:
	@echo "Repairing..."
	@supabase migration repair --status applied
	@echo "Repaired!"

docker-stop-all:
	@echo "Stopping all containers..."
	@docker stop $$(docker ps -a -q)
	@echo "Stopped all containers!"

.PHONY: all build run test clean
