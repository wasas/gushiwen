package poet

import (
	"github.com/trytwice/gushiwen/model"
	"github.com/trytwice/gushiwen/pkg/db"
)

func countPoet() (float64, error) {
	db, err := db.OpenDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	count := model.Poet{}
	db.Last(&count)
	return float64(count.ID), nil
}
