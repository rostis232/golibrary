package postgres

import (
	"fmt"

	"github.com/rostis232/golibrary/models"
)

func (p *Postgres) getAllDifficulties() (map[int64]models.Difficulty, error) {
	diffMap := make(map[int64]models.Difficulty)
	difficulties := []models.Difficulty{}
	query := fmt.Sprintf("select * from %s", difficultyTable)
	err := p.db.Select(&difficulties, query)
	if err != nil {
		return diffMap, fmt.Errorf("error occured in GetAllDifficulties %w", err)
	}
	for _, difficulty := range difficulties {
		diffMap[difficulty.ID] = difficulty
	}
	return diffMap, nil
}
