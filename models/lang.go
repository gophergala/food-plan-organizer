package models

import "database/sql"

type Language struct {
	NutrientID string
	FactorCode string
}

func InsertLanguage(l *Language, tx *sql.DB) error {
	if _, err := tx.Exec(`INSERT INTO languages VALUES ($1, $2);`, l.NutrientID, l.FactorCode); err != nil {
		return err
	}
	return nil
}
