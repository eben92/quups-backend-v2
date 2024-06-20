#!/bin/bash

# Function to install goose if not already installed
install_goose() {
    if ! command -v goose &> /dev/null; then
        echo "goose not found, installing..."
        go get -u github.com/pressly/goose/v3/cmd/goose
        if [ $? -ne 0 ]; then
            echo "Failed to install goose. Please install manually and try again."
            exit 1
        fi
    fi
}

# Install goose if not installed
install_goose

# Set database URL and run goose migrations
DATABASE_URL="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable"
GOOSE_MIGRATION_DIR="./internal/database/migrations"
GOOSE_DRIVER="postgres"

echo "Running goose migrations..."
goose -dir "${GOOSE_MIGRATION_DIR}" -driver "${GOOSE_DRIVER}" -url "${DATABASE_URL}" up
