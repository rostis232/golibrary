package postgres

import (
	"fmt"

	"github.com/rostis232/golibrary/models"
)

func(p *Postgres) getAllLanguages()(map[int64]models.Language, error){
	langMap := make(map[int64]models.Language)
	query := fmt.Sprintf("select * from %s", languageTable)
	err := p.db.Select(&langMap, query)
	if err != nil {
		return langMap, fmt.Errorf("error occured in GetAllLanguages %w", err)
	}
	return langMap, nil
}