package postgres

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"

)

const (
	libraryTable = "library"
)

func NewPostgres(configDB string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", configDB)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}