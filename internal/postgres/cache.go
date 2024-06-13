package postgres

import (
	"fmt"
)

func(p *Postgres) UpdateCachedTypes() error {
	typesMap := map[int64]map[string]string{}
	types, err := p.getAllTypes()
	if err != nil {
		return fmt.Errorf("error getting cache for types: %w", err)
	}
	for _, t := range types {
		typeMap := map[string]string{}
		typeMap["UK"] = t.TitleUK
		typeMap["EN"] = t.TitleEN
		typesMap[t.ID] = typeMap
	}
	p.CachedTypes = typesMap
	return nil
}

func(p *Postgres) UpdateCachedLanguages() error {
	languagesMap := map[int64]map[string]string{}
	languages, err := p.getAllLanguages()
	if err != nil {
		return fmt.Errorf("error getting cache for languages: %w", err)
	}
	for _, l := range languages{
		languageMap := map[string]string{}
		languageMap["UK"] = l.TitleUK
		languageMap["EN"] = l.TitleEN
		languagesMap[l.ID] = languageMap
	}
	p.CachedLanguages = languagesMap
	return nil
}

func(p *Postgres) UpdateCachedDifficulties() error {
	difficultiesMap := map[int64]map[string]string{}
	difficulties, err := p.getAllDifficulties()
	if err != nil {
		return fmt.Errorf("error getting cache for difficulties: %w", err)
	}
	for _, d := range difficulties{
		difficultyMap := map[string]string{}
		difficultyMap["UK"] = d.TitleUK
		difficultyMap["EN"] = d.TitleEN
		difficultiesMap[d.ID] = difficultyMap
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

func(p *Postgres) GetCachedTypes() map[int64]map[string]string {
	return p.CachedTypes
}

func(p *Postgres) GetCachedLanguages() map[int64]map[string]string {
	return p.CachedLanguages
}

func(p *Postgres) GetCachedDifficulties() map[int64]map[string]string {
	return p.CachedDifficulties
}