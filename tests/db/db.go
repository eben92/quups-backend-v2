package testdb

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type Queries interface {
}

type TestDatabase struct {
	container testcontainers.Container
}

func NewDatabase(t *testing.T) *TestDatabase {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:12",
		ExposedPorts: []string{"5432"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "postgres",
		},

		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	require.NoError(t, err)

	return &TestDatabase{
		container: container,
	}
}

func (db *TestDatabase) Port(t *testing.T) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	p, err := db.container.MappedPort(ctx, "5432")

	require.NoError(t, err)
	return p.Int()
}

func (db *TestDatabase) ConnectionString(t *testing.T) string {
	return fmt.Sprintf("postgres://postgres:postgres@127.0.0.1:%d/postgres", db.Port(t))
}

func (db *TestDatabase) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	require.NoError(t, db.container.Terminate(ctx))
}
