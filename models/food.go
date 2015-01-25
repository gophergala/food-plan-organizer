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
