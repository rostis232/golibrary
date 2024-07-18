package postgres

import (
	"fmt"
)

func (p *Postgres) UpdateCachedTypes() error {
	typesMap := map[int64]string{}
	types, err := p.getAllTypes()
	if err != nil {
		return fmt.Errorf("error getting cache for types: %w", err)
	}
	for _, t := range types {
		typesMap[t.ID] = t.Title
	}
	p.CachedTypes = typesMap
	return nil
}

func (p *Postgres) UpdateCachedLanguages() error {
	languagesMap := map[int64]string{}
	languages, err := p.getAllLanguages()
	if err != nil {
		return fmt.Errorf("error getting cache for languages: %w", err)
	}
	for _, l := range languages {
		languagesMap[l.ID] = l.Title
	}
	p.CachedLanguages = languagesMap
	return nil
}

func (p *Postgres) UpdateCachedDifficulties() error {
	difficultiesMap := map[int64]string{}
	difficulties, err := p.getAllDifficulties()
	if err != nil {
		return fmt.Errorf("error getting cache for difficulties: %w", err)
	}
	for _, d := range difficulties {
		difficultiesMap[d.ID] = d.Title
	}
	p.CachedDifficulties = difficultiesMap
	return nil
}

func (p *Postgres) InitCache() error {
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

func (p *Postgres) GetCachedTypes() map[int64]string {
	return p.CachedTypes
}

func (p *Postgres) GetCachedLanguages() map[int64]string {
	return p.CachedLanguages
}

func (p *Postgres) GetCachedDifficulties() map[int64]string {
	return p.CachedDifficulties
}
