package postgres

import (
	"fmt"

	"github.com/rostis232/golibrary/models"
)

func (p *Postgres) getAllTypes() (map[int64]models.Type, error) {
	typesMap := make(map[int64]models.Type)
	types := []models.Type{}
	query := fmt.Sprintf("select * from %s", typeTable)
	err := p.db.Select(&types, query)
	if err != nil {
		return typesMap, fmt.Errorf("error occured in GetAllTypes %w", err)
	}
	for _, t := range types {
		typesMap[t.ID] = t
	}
	return typesMap, nil
}
