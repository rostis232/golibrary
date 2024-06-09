package postgres

import "github.com/jmoiron/sqlx"

type Library struct{
	db *sqlx.DB
}

func NewLibrary(db *sqlx.DB) *Library {
	return &Library{db: db}
}

func(l *Library) GetAllItems(){}