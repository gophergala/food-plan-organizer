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
}

var CreateNutrientTableSQLs = []string{`
  CREATE TABLE nutrients (
    food_id            string,
    nutrient_id        string,
    nutrient_value     float,
    min                float,
    max                float,
    degrees_of_freedom int,
    lower_error_bound  float,
    upper_error_bound  float
  );`,
	`CREATE INDEX nutrients_food_idx ON nutrients (food_id);`,
	`CREATE INDEX nutrients_nutrient_idx ON nutrients (nutrient_id);`,
}

func InsertNutrient(n *Nutrient, tx *sql.Tx) error {
	if _, err := tx.Exec(`INSERT INTO nutrients VALUES ($1,$2,$3,$4,$5,$6,$7,$8);`, n.FoodID, n.NutrientID, n.NutritionValue, n.Min, n.Max, n.DegreesOfFreedom, n.LowerErrorBound, n.UpperErrorBound); err != nil {
		return err
	}
	return nil
}
