package models

import "database/sql"

type FoodGroup struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

var CreateFoodGroupTableSQLs = []string{`
  CREATE TABLE food_groups (
    id   integer,
    name text
  );`,
	`CREATE INDEX food_groups_idx ON food_groups (id);`,
}

func InsertFoodGroup(fg *FoodGroup, tx *sql.DB) error {
	if _, err := tx.Exec(`INSERT INTO food_groups VALUES ($1,$2);`, fg.ID, fg.Name); err != nil {
		return err
	}
	return nil
}
