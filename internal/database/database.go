package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"

	"quups-backend/internal/database/repository"
)

type Service interface {
	Health() map[string]string
	NewRepository() *repository.Queries
	NewRawDB() *sql.DB
	CreateTable(tableName string, columns []string) error
}

type service struct {
	db         *sql.DB
	repository *repository.Queries
}

var (
	dbInstance *service
)

func NewService(connStr string) Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

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

func (s *service) CreateTable(tableName string, columns []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cols := ""
	for _, col := range columns {

		if col == columns[len(columns)-1] {
			cols += col
		} else {
			cols += col + ", "
		}
	}

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, cols)
	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	slog.Info("Table created successfully", "TABLE", tableName)

	return nil
}
