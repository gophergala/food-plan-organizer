package models

import "database/sql"

type NutrientDefinition struct {
	NutrientID    int32  `json:"id"`
	Units         string `json:"unit"`
	Tagname       string `json:"tagname"`
	Description   string `json:"description"`
	DecimalPlaces int32  `json:"decimal_places"`
}

var CreateNutrientDefinitionTableSQLs = []string{`
  CREATE TABLE nutrient_definitions (
    nutrient_id    integer,
    units          TEXT,
    tagname        TEXT,
    description    TEXT,
    decimal_places integer
  );`,
	`CREATE INDEX nutrient_definitions_idx ON nutrient_definitions (nutrient_id);`,
}

func InsertNutrientDefinition(nd *NutrientDefinition, tx *sql.DB) error {
	if _, err := tx.Exec(`INSERT INTO nutrient_definitions VALUES ($1,$2,$3,$4,$5);`, nd.NutrientID, nd.Units, nd.Tagname, nd.Description, nd.DecimalPlaces); err != nil {
		return err
	}
	return nil
}
