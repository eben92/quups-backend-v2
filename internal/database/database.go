package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"

	"quups-backend/internal/database/repository"
)

type Service interface {
	Health() map[string]string
	NewRepository() *repository.Queries
	NewRawDB() *sql.DB
}

type service struct {
	db         *sql.DB
	repository *repository.Queries
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	dbInstance *service
)

func NewService() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		username,
		password,
		host,
		port,
		database,
	)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &service{
		db:         db,
		repository: repository.New(db),
	}
	return dbInstance
}

func (s *service) NewRepository() *repository.Queries {
	return s.repository
}

func (s *service) NewRawDB() *sql.DB {
	return s.db
}

type Repo struct {
	Queries *repository.Queries
	DB      *sql.DB
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
