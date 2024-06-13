package postgres

import (
	"fmt"

	"github.com/rostis232/golibrary/models"
)

func (p *Postgres) GetAllLibraryItems() ([]models.LibraryItem, error) {
	items := []models.LibraryItem{}
	query := fmt.Sprintf("select * from %s", libraryTable)
	err := p.db.Select(&items, query)
	if err != nil {
		return items, fmt.Errorf("error occured in GetAllLibraryItems %w", err)
	}
	return items, nil
}
