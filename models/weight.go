package models

import "database/sql"

type Weight struct {
	NutrientID  int32
	Seq         string
	Amount      float32
	Description string
	GramWeight  float32
}

var CreateWeightTableSQLs = []string{`
  CREATE TABLE weights (
    nutrient_id integer,
    seq         text,
    amount      real,
    description text,
    gram_weight real
  );`,
	`CREATE INDEX weights_idx ON weights (nutrient_id);`,
}

func InsertWeight(w *Weight, tx *sql.DB) error {
	if _, err := tx.Exec(`INSERT INTO weights VALUES ($1, $2, $3, $4, $5);`, w.NutrientID, w.Seq, w.Amount, w.Description, w.GramWeight); err != nil {
		return err
	}
	return nil
}
