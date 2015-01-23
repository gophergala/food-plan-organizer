package models

import "database/sql"

type Nutrient struct {
	FoodID           string
	NutrientID       string
	NutritionValue   float32
	Min              float32
	Max              float32
	DegreesOfFreedom int32
	LowerErrorBound  float32
	UpperErrorBound  float32
	NutrientDefinition
}

var CreateNutrientTableSQLs = []string{`
  CREATE TABLE nutrients (
    food_id            TEXT,
    nutrient_id        TEXT,
    nutrient_value     REAL,
    min                REAL,
    max                REAL,
    degrees_of_freedom INTEGER,
    lower_error_bound  REAL,
    upper_error_bound  REAL
  );`,
	`CREATE INDEX nutrients_food_idx ON nutrients (food_id);`,
	`CREATE INDEX nutrients_nutrient_idx ON nutrients (nutrient_id);`,
}

// ScanExtendedNutrient scans nutrient, then nutrient_definition
func ScanNutrient(s StructScanner) (Nutrient, error) {
	var n Nutrient
	var err = s.Scan(&n.FoodID, &n.NutrientID, &n.NutritionValue, &n.Min, &n.Max, &n.DegreesOfFreedom, &n.LowerErrorBound, &n.UpperErrorBound, &n.NutrientDefinition.NutrientID, &n.Units, &n.Tagname, &n.Description, &n.DecimalPlaces)
	if err != nil {
		return n, err
	}
	return n, nil
}

func InsertNutrient(n *Nutrient, tx *sql.DB) error {
	if _, err := tx.Exec(`INSERT INTO nutrients VALUES ($1,$2,$3,$4,$5,$6,$7,$8);`, n.FoodID, n.NutrientID, n.NutritionValue, n.Min, n.Max, n.DegreesOfFreedom, n.LowerErrorBound, n.UpperErrorBound); err != nil {
		return err
	}
	return nil
}
