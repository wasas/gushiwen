package poetry

import (
	"github.com/trytwice/gushiwen/model"
	"github.com/trytwice/gushiwen/pkg/db"
)

func countPoetry() (float64, error) {
	db, err := db.OpenDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	count := model.Poetry{}
	db.Last(&count)
	return float64(count.ID), nil
}
