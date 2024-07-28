package testdb_test

import (
	testdb "quups-backend/tests/db"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	db := testdb.NewDatabase(t)
	defer db.Close(t)
	connStr := db.ConnectionString(t)
	println(connStr)
}
