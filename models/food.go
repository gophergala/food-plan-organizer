package models

import "database/sql"

type Food struct {
	ID                  string
	FoodGroupID         string
	Name                string
	ShortName           string
	CommonName          string
	ScientificName      string
	NitrogenFactor      float32
	ProteinFactor       float32
	FatFactor           float32
	CarbonhydrateFactor float32
}

var CreateFoodTableSQLs = []string{`
  CREATE TABLE foods (
    id                   string,
    food_group_id        string,
    name                 string,
    short_name           string,
    common_name          string,
    scientific_name      string,
    nitrogen_factor      float,
    protein_factor       float,
    fat_factor           float,
    carbonhydrate_factor float
  );`,
	`CREATE INDEX foods_idx ON foods (id);`,
	`CREATE INDEX foods_food_group_idx ON foods (food_group_id);`,
}

func ScanFood(s StructScanner) (Food, error) {
	var f Food
	var err = s.Scan(&f.ID, &f.FoodGroupID, &f.Name, &f.ShortName, &f.CommonName, &f.ScientificName, &f.NitrogenFactor, &f.ProteinFactor, &f.FatFactor, &f.CarbonhydrateFactor)
	if err != nil {
		return f, err
	}
	return f, nil
}

func InsertFood(f *Food, tx *sql.Tx) error {
	if _, err := tx.Exec(`INSERT INTO foods VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);`, f.ID, f.FoodGroupID, f.Name, f.ShortName, f.CommonName, f.ScientificName, f.NitrogenFactor, f.ProteinFactor, f.FatFactor, f.CarbonhydrateFactor); err != nil {
		return err
	}
	return nil
}
