package testdb_test

import (
	testdb "quups-backend/tests/db"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	db := testdb.NewTestDatabase(t)
	defer db.Close(t)
	connStr := db.ConnectionString(t)
	println(connStr)
}
