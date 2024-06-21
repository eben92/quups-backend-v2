#!/bin/bash

# Set database URL and run goose migrations
DATABASE_URL="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable"
GOOSE_MIGRATION_DIR="./internal/database/migrations"
GOOSE_DRIVER="postgres"

echo "Running goose migrations..."
goose -dir "${GOOSE_MIGRATION_DIR}" ${GOOSE_DRIVER} "host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_DATABASE} sslmode=disable" up
echo "Migrated!"
