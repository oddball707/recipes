package migrations

import (
	"database/sql"
	"embed"

	_ "github.com/jackc/pgx/v5/stdlib" // pgx driver
	"github.com/pressly/goose/v3"
)

//go:embed sql
var directory embed.FS

func Up(db *sql.DB) error {
	goose.SetBaseFS(directory)
	goose.SetTableName("goose_db_version")

	return goose.Up(db, "sql", goose.WithAllowMissing())
}
