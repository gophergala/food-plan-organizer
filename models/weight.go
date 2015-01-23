package models

import "database/sql"

type Weight struct {
	NutrientID  string
	Seq         string
	Amount      float32
	Description string
	GramWeight  float32
}

var CreateWeightTableSQLs = []string{`
  CREATE TABLE weights (
    nutrient_id string,
    seq         string,
    amount      float,
    description string,
    gram_weight float
  );`,
	`CREATE INDEX weights_idx ON weights (nutrient_id);`,
}

func InsertWeight(w *Weight, tx *sql.Tx) error {
	if _, err := tx.Exec(`INSERT INTO weights VALUES ($1, $2, $3, $4, $5);`, w.NutrientID, w.Seq, w.Amount, w.Description, w.GramWeight); err != nil {
		return err
	}
	return nil
}
