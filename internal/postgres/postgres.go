package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rostis232/golibrary/models"
)

const (
	libraryTable = "library_item"
	languageTable = "language"
	difficultyTable = "difficulty"
	typeTable = "type"
)

type Postgres struct {
	db *sqlx.DB
	CachedTypes map[int64]models.Type
	CachedLanguages map[int64]models.Language
	CachedDifficulties map[int64]models.Difficulty
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
