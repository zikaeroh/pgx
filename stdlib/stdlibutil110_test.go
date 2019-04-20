// +build go1.10

package stdlib_test

import (
	"database/sql"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

func openDB(t *testing.T) *sql.DB {
	config, err := pgx.ParseConfig("postgres://pgx_md5:secret@127.0.0.1:5432/pgx_test")
	if err != nil {
		t.Fatalf("pgx.ParseConnectionString failed: %v", err)
	}

	return stdlib.OpenDB(*config)
}
