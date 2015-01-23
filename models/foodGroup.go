package models

import "database/sql"

type FoodGroup struct {
	ID   string
	Name string
}

var CreateFoodGroupTableSQLs = []string{`
  CREATE TABLE food_groups (
    id   string,
    name string
  );`,
	`CREATE INDEX food_groups_idx ON food_groups (id);`,
}

func InsertFoodGroup(fg *FoodGroup, tx *sql.Tx) error {
	if _, err := tx.Exec(`INSERT INTO food_groups VALUES ($1, $2);`, fg.ID, fg.Name); err != nil {
		return err
	}
	return nil
}