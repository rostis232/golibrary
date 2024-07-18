package postgres

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

const (
	libraryTable    = "library_item"
	languageTable   = "language"
	difficultyTable = "difficulty"
	typeTable       = "type"
)

type Postgres struct {
	db                 *sqlx.DB
	CachedTypes        map[int64]string
	CachedLanguages    map[int64]string
	CachedDifficulties map[int64]string
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

	p := Postgres{db: db}

	err = p.Migrate()
	if err != nil {
		return nil, err
	}

	err = p.InitCache()
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *Postgres) Migrate() error {
	log.Infof("migrating database")
	driver, err := postgres.WithInstance(p.db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("postgres: could not instantiate database driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./schema",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("postgres: could not instantiate migrate instance: %w", err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("postgres: could not run migrations: %w", err)
	}
	log.Infof("migrated database")
	return nil
}
