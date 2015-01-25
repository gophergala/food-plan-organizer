package models

import "database/sql"

type LanguageDescription struct {
	Code        string
	Description string
}

func InsertLanguageDescription(ld *LanguageDescription, tx *sql.DB) error {
	if _, err := tx.Exec(`INSERT INTO language_descriptions VALUES ($1, $2);`, ld.Code, ld.Description); err != nil {
		return err
	}
	return nil
}
