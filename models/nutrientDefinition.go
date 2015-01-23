package models

import "database/sql"

type NutrientDefinition struct {
	NutrientID    string
	Units         string
	Tagname       string
	Description   string
	DecimalPlaces int32
}

var CreateNutrientDefinitionTableSQLs = []string{`
  CREATE TABLE nutrient_definitions (
    nutrient_id    string,
    units          string,
    tagname        string,
    description    string,
    decimal_places int
  );`,
	`CREATE INDEX nutrient_definitions_idx ON nutrient_definitions (nutrient_id);`,
}

func InsertNutrientDefinition(nd *NutrientDefinition, tx *sql.Tx) error {
	if _, err := tx.Exec(`INSERT INTO nutrient_definitions VALUES ($1,$2,$3,$4,$5);`, nd.NutrientID, nd.Units, nd.Tagname, nd.Description, nd.DecimalPlaces); err != nil {
		return err
	}
	return nil
}
