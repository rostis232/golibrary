package postgres

import (
	"fmt"

	"github.com/rostis232/golibrary/models"
)

func(p *Postgres) getAllLanguages()(map[int64]models.Language, error){
	langMap := make(map[int64]models.Language)
	languages := []models.Language{}
	query := fmt.Sprintf("select * from %s", languageTable)
	err := p.db.Select(&languages, query)
	if err != nil {
		return langMap, fmt.Errorf("error occured in GetAllLanguages %w", err)
	}
	for _, language := range languages {
		langMap[language.ID] = language
	}
	return langMap, nil
}