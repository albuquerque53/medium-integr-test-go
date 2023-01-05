package db

import (
	"context"
	"database/sql"
)

type DBManager struct {
	DB *sql.DB
}

func NewDBManager(DB *sql.DB) *DBManager {
	return &DBManager{
		DB: DB,
	}
}

func (dm *DBManager) Query(ctx context.Context, q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := dm.DB.Prepare(q)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.QueryContext(ctx, args...)
}

func (dm *DBManager) QueryRow(ctx context.Context, q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := dm.DB.Prepare(q)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...), nil
}
