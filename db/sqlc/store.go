package db

import "database/sql"

type StoreExec struct {
	db *sql.DB
	*Queries
}

func NewStoreExec(db *sql.DB) *StoreExec {
	return &StoreExec{
		db:      db,
		Queries: New(db),
	}
}