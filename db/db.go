package db

import (
	"database/sql"
	"log/slog"

	_ "modernc.org/sqlite"
)

func NewSQLLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "file:data.db")
	if err != nil {
		slog.Error("sql open error", slog.Any("error", err))
		return nil, err
	}
	if err := RunMigration(db); err != nil {
		slog.Error("run migration error", slog.Any("error", err))
		return nil, err
	}
	return db, nil
}
