package postgres

import (
	"fmt"

	"github.com/rostis232/golibrary/models"
)

func(p *Postgres) UpdateCachedTypes() error {
	typesMap := map[int64]models.Type{}
	types, err := p.getAllTypes()
	if err != nil {
		return fmt.Errorf("error getting cache for types: %w", err)
	}
	for _, t := range types {
		typesMap[t.ID] = t
	}
	p.CachedTypes = typesMap
	return nil
}

func(p *Postgres) UpdateCachedLanguages() error {
	languagesMap := map[int64]models.Language{}
	languages, err := p.getAllLanguages()
	if err != nil {
		return fmt.Errorf("error getting cache for languages: %w", err)
	}
	for _, l := range languages{
		languagesMap[l.ID] = l
	}
	p.CachedLanguages = languagesMap
	return nil
}

func(p *Postgres) UpdateCachedDifficulties() error {
	difficultiesMap := map[int64]models.Difficulty{}
	difficulties, err := p.getAllDifficulties()
	if err != nil {
		return fmt.Errorf("error getting cache for difficulties: %w", err)
	}
	for _, d := range difficulties{
		difficultiesMap[d.ID] = d
	}
	p.CachedDifficulties = difficultiesMap
	return nil
}

func(p *Postgres) InitCache() error {
	if err := p.UpdateCachedDifficulties(); err != nil {
		return err
	}
	if err := p.UpdateCachedLanguages(); err != nil {
		return err
	}
	if err := p.UpdateCachedTypes(); err != nil {
		return err
	}
	return nil
}