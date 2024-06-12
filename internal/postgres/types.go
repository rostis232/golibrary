package postgres

import (
	"fmt"

	"github.com/rostis232/golibrary/models"
)

func(p *Postgres) getAllTypes()(map[int64]models.Type, error){
	typesMap := make(map[int64]models.Type)
	query := fmt.Sprintf("select * from %s", typeTable)
	err := p.db.Select(&typesMap, query)
	if err != nil {
		return typesMap, fmt.Errorf("error occured in GetAllTypes %w", err)
	}
	return typesMap, nil
}