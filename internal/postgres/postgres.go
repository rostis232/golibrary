package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	libraryTable = "library_item"
	languageTable = "language"
	difficultyTable = "difficulty"
	typeTable = "type"
)

type Postgres struct {
	db *sqlx.DB
	CachedTypes map[int64]map[string]string
	CachedLanguages map[int64]map[string]string
	CachedDifficulties map[int64]map[string]string
}

func NewPostgres(configDB string) (*Postgres, error) {
	db, err := sqlx.Open("postgres", configDB)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	p := Postgres{db:db}

	err = p.InitCache()
	if err != nil {
		return nil, err
	}

	return &p, nil
}
