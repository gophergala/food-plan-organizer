package models

import "database/sql"

type Food struct {
	ID                  int32      `json:"id"`
	FoodGroupID         int32      `json:"-"`
	Group               FoodGroup  `json:"food_group"`
	Name                string     `json:"name"`
	ShortName           string     `json:"short_name"`
	CommonName          string     `json:"common_name"`
	ScientificName      string     `json:"scientific_name"`
	NitrogenFactor      float32    `json:"nitrogen_factor"`
	ProteinFactor       float32    `json:"protein_factor"`
	FatFactor           float32    `json:"fat_factor"`
	CarbonhydrateFactor float32    `json:"carbonhydrate_factor"`
	Nutrients           []Nutrient `json:"nutrients,omitempty"`
}

var CreateFoodTableSQLs = []string{`
  CREATE TABLE foods (
    id                   integer,
    food_group_id        integer,
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
	var err = s.Scan(&f.ID, &f.FoodGroupID, &f.Name, &f.ShortName, &f.CommonName, &f.ScientificName, &f.NitrogenFactor, &f.ProteinFactor, &f.FatFactor, &f.CarbonhydrateFactor, &f.Group.ID, &f.Group.Name)
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
