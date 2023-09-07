package postgres

import (
	"database/sql"
)

type storagePg struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *storagePg {
	return &storagePg{
		db: db,
	}
}
