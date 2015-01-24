package models

type Recipe struct {
	ID          int32        `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Ingredients []Ingredient `json:"ingredients"`
}
type Ingredient struct {
	ID                  int32      `json:"-"`
	RecipeID            int32      `json:"-"`
	FoodID              int32      `json:"food_id"`
	Unit                string     `json:"unit"`
	Volume              float32    `json:"volume"`
	Name                string     `json:"name"`
	NitrogenFactor      float32    `json:"nitrogen_factor"`
	ProteinFactor       float32    `json:"protein_factor"`
	FatFactor           float32    `json:"fat_factor"`
	CarbonhydrateFactor float32    `json:"carbonhydrate_factor"`
	Nutrients           []Nutrient `json:"nutrients"`
}

var CreateRecipeTableSQLs = []string{`
  CREATE TABLE IF NOT EXISTS recipes (
    id             integer PRIMARY KEY ASC,
    name           TEXT,
    description    TEXT
  );`,
	// `DROP TABLE ingredients;`,
	`CREATE TABLE IF NOT EXISTS ingredients (
    id          integer primary key asc,
    recipe_id   integer,
    food_id     integer,
    unit        text,
    volume      real
  );`,
}
