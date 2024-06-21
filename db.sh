#!/bin/bash

# Function to install go if not already installed
install_go(){
 if command -v go >/dev/null 2>&1; then
    echo "Go is installed. Proceeding with script..."
    go version 
else
    echo "Go is not installed. Installing go..."

    curl -sSL https://go.dev/dl/go1.22.4.linux-amd64.tar.gz | tar -C /usr/local/ -xz
    export PATH:$PATH:/usr/local/go/bin

    #Verify installation
    go version 
 fi
}

# Function to install goose if not already installed
install_goose() {
    if ! command -v goose >/dev/null 2>&1; then
        echo "goose not found, installing..."
        go get -u github.com/pressly/goose/v3/cmd/goose@latest
        if [ $? -ne 0 ]; then
            echo "Failed to install goose. Please install manually and try again."
            exit 1
        fi
    fi
}

# Install go if it doesnt exist before installing goose
#install_go

# Install goose if not installed
#install_goose

# Set database URL and run goose migrations
DATABASE_URL="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable"
GOOSE_MIGRATION_DIR="./internal/database/migrations"
GOOSE_DRIVER="postgres"

echo "Running goose migrations..."
goose -dir "${GOOSE_MIGRATION_DIR}" ${GOOSE_DRIVER} "host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_DATABASE} sslmode=disable" up
echo "Migrated!"
