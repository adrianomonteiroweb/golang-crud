package db

import "database/sql"

type StoreExec struct {
	db *sql.DB
	*Queries
}

func ExecuteNewStore(db *sql.DB) *StoreExec {
	return &StoreExec{
		db:      db,
		Queries: New(db),
	}
}