package testdb

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	utils_test "quups-backend/tests/utils"
	"runtime"
	"testing"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestDatabase struct {
	container testcontainers.Container
}

func NewTestDatabase(t *testing.T) *TestDatabase {

	slog.Info("Creating new database instance")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	net, err := network.New(ctx)
	require.NoError(t, err, "Failed to create Docker network")

	postgresCtn := createPostgresContainer(ctx, t, net)
	require.NoError(t, utils_test.WaitForContainerReady(ctx, postgresCtn), "Postgres container is not ready")

	gooseCtn := runMigration(ctx, t, net)
	defer gooseCtn.Terminate(ctx)
	require.NoError(t, utils_test.CheckContainerConnectivity(ctx, gooseCtn, "postgres", "5432"), "Goose cannot connect to Postgres")

	return &TestDatabase{
		container: postgresCtn,
	}

}

func createPostgresContainer(ctx context.Context, t *testing.T, network *testcontainers.DockerNetwork) testcontainers.Container {
	_, path, _, ok := runtime.Caller(0)
	sqlfiles, err := filepath.Glob(filepath.Join(filepath.Dir(path), "..", "..", "internal", "database", "migrations", "*.sql"))

	if !ok {
		require.NoError(t, err, "failed to get path")
	}

	creq := testcontainers.ContainerRequest{
		Image:        "postgres:12",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "postgres",
		},
		Networks:       []string{network.Name},
		NetworkAliases: map[string][]string{network.Name: {"postgres"}},
		WaitingFor:     wait.ForListeningPort("5432/tcp"),
	}

	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: creq,
		Started:          true,
	})

	require.NoError(t, err, "could not start postgres container")

	fmt.Println("total sql files", len(sqlfiles))

	return c

}

func runMigration(ctx context.Context, t *testing.T, network *testcontainers.DockerNetwork) testcontainers.Container {

	_, path, _, ok := runtime.Caller(0)
	require.True(t, ok)
	absPath, err := filepath.Abs(filepath.Join(filepath.Dir(path), "..", "..", "internal", "database", "migrations"))

	require.NoError(t, err)

	creq := testcontainers.ContainerRequest{
		Image: "ghcr.io/eben92/goose-docker:latest",
		Env: map[string]string{
			"GOOSE_DRIVER":   "postgres",
			"GOOSE_DBSTRING": "host=postgres port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
			"GOOSE_COMMAND":  "up",
		},
		Networks:       []string{network.Name},
		NetworkAliases: map[string][]string{network.Name: {"goose"}},
		HostConfigModifier: func(h *container.HostConfig) {
			h.Mounts = []mount.Mount{
				{
					Type:     mount.TypeBind,
					Source:   absPath,
					Target:   "/migrations",
					ReadOnly: false,
					BindOptions: &mount.BindOptions{
						Propagation: "rprivate",
					},
				},
			}
		},
	}

	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: creq,
		Started:          true,
	})
	require.NoError(t, err, "could not start goose container")

	return c
}

func (db *TestDatabase) Port(t *testing.T) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	p, err := db.container.MappedPort(ctx, "5432")

	require.NoError(t, err)

	return p.Int()
}

func (db *TestDatabase) ConnectionString(t *testing.T) string {
	return fmt.Sprintf("postgres://postgres:postgres@127.0.0.1:%d/postgres?sslmode=disable", db.Port(t))
}

func (db *TestDatabase) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	require.NoError(t, db.container.Terminate(ctx))
}
