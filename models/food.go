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
	Nutrients           []Nutrient
}

var CreateFoodTableSQLs = []string{`
  CREATE TABLE foods (
    id                   text,
    food_group_id        text,
    name                 text,
    short_name           text,
    common_name          text,
    scientific_name      text,
    nitrogen_factor      real,
    protein_factor       real,
    fat_factor           real,
    carbonhydrate_factor real
  );`,
	`CREATE INDEX foods_idx ON foods (id);`,
	`CREATE INDEX foods_food_group_idx ON foods (food_group_id);`,
	`CREATE INDEX foods_name ON foods (name);`,
	`CREATE INDEX foods_short_name ON foods (short_name);`,
	`CREATE INDEX foods_common_name ON foods (common_name);`,
	`CREATE INDEX foods_scientific_name ON foods (scientific_name);`,
}

func ScanFood(s StructScanner) (Food, error) {
	var f Food
	var err = s.Scan(&f.ID, &f.FoodGroupID, &f.Name, &f.ShortName, &f.CommonName, &f.ScientificName, &f.NitrogenFactor, &f.ProteinFactor, &f.FatFactor, &f.CarbonhydrateFactor)
	if err != nil {
		return f, err
	}
	return f, nil
}

func InsertFood(f *Food, tx *sql.DB) error {
	if _, err := tx.Exec(`INSERT INTO foods VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);`, f.ID, f.FoodGroupID, f.Name, f.ShortName, f.CommonName, f.ScientificName, f.NitrogenFactor, f.ProteinFactor, f.FatFactor, f.CarbonhydrateFactor); err != nil {
		return err
	}
	return nil
}
