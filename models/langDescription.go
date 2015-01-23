package models

import "database/sql"

type LanguageDescription struct {
	Code        string
	Description string
}

var CreateLanguageDescriptionTableSQLs = []string{`
  CREATE TABLE language_descriptions (
    factor_code string,
    description string
  );`,
	`CREATE INDEX language_descriptions_idx ON language_descriptions (factor_code);`,
}

func InsertLanguageDescription(ld *LanguageDescription, tx *sql.Tx) error {
	if _, err := tx.Exec(`INSERT INTO language_descriptions VALUES ($1, $2);`, ld.Code, ld.Description); err != nil {
		return err
	}
	return nil
}
