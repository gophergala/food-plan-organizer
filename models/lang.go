package models

import "database/sql"

type Language struct {
	NutrientID string
	FactorCode string
}

var CreateLanguageTableSQLs = []string{`
  CREATE TABLE languages (
    nutrient_id string,
    factor_code string
  );`,
	`CREATE INDEX languages_nutrient_idx ON languages (nutrient_id);`,
	`CREATE INDEX languages_factor_code_idx ON languages (factor_code);`,
}

func InsertLanguage(l *Language, tx *sql.Tx) error {
	if _, err := tx.Exec(`INSERT INTO languages VALUES ($1, $2);`, l.NutrientID, l.FactorCode); err != nil {
		return err
	}
	return nil
}
