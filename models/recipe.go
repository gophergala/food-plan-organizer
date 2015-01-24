package models

import "database/sql"

type Recipe struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Ingredient struct {
	ID     int32  `json:"-"`
	Name   string `json:"name"`
	Unit   string `json:"unit"`
	FoodID int32  `json:"food_id"`
}
type IngredientUsage struct {
	IngredientID int32
	RecipeID     int32
	Volume       float32
}

var CreateRecipeTableSQLs = []string{`
  CREATE TABLE IF NOT EXISTS recipes (
    id             integer PRIMARY KEY ASC,
    name           TEXT,
    description    TEXT
  );`,
	// `DROP TABLE ingredients;`,
	`CREATE TABLE IF NOT EXISTS ingredients (
    id       integer primary key asc,
    name     text,
    food_id  integer,
    unit     text
  );`,
	`CREATE TABLE IF NOT EXISTS ingredient_usages (
    ingridient_id integer,
    recipe_id     integer,
    volume        real
  );`,
}

func InsertRecipe(r *Recipe, tx *sql.DB) error {
	if _, err := tx.Exec(`INSERT INTO recipes VALUES ($1,$2,$3);`, r.ID, r.Name, r.Description); err != nil {
		return err
	}
	return nil
}
