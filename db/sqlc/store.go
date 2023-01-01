package db

import (
	"database/sql"
)



type Store struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) *Store {
    store := Store{
		db: db,
		Queries: New(db),
	}
	return &store
}