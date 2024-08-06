package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"quups-backend/internal/database"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	db   database.Service
}

var (
	dbname   = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	dbport   = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func NewServer() *http.Server {

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		username,
		password,
		host,
		dbport,
		dbname,
	)

	db := database.NewService(connStr)
	NewServer := &Server{
		port: port,
		db:   db,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	slog.Info("Server started on port", "APP", port)

	return server
}
