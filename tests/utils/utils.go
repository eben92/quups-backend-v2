package utils_test

import (
	"bufio"
	"context"
	"fmt"
	"strings"

	"github.com/testcontainers/testcontainers-go"
)

// checkContainerConnectivity  checks if containers can communicate with each other
func CheckContainerConnectivity(ctx context.Context, container testcontainers.Container, targetHost string, targetPort string) error {
	cmd := []string{"nc", "-zv", targetHost, targetPort}

	exitCode, r, err := container.Exec(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to execute command in container: %w", err)
	}
	if exitCode != 0 {
		return fmt.Errorf("container could not connect to %s:%s", targetHost, targetPort)
	}

	scanner := bufio.NewScanner(r)

	c, _ := container.Inspect(ctx)

	for scanner.Scan() {
		fmt.Println("net info ==>", c.Name, scanner.Text())
	}

	return nil
}

func WaitForContainerReady(ctx context.Context, container testcontainers.Container) error {
	expectedLogMessage := "started"

	logs, err := container.Logs(ctx)
	if err != nil {
		return fmt.Errorf("failed to get logs from container: %w", err)
	}
	defer logs.Close()

	scanner := bufio.NewScanner(logs)
	for scanner.Scan() {

		if strings.Contains(scanner.Text(), expectedLogMessage) {
			return nil // Container is ready
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading container logs: %w", err)
	}

	return fmt.Errorf("container did not become ready within the timeout period")
}
