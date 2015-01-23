package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cznic/ql"
	"github.com/nicolai86/sr27/models"
)

func main() {
	ql.RegisterDriver()
	var db, err = sql.Open("ql", "test.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v\n", err)
	}
	defer db.Close()

	var rows *sql.Rows
	// rows, err = db.Query(`SELECT Name, Schema FROM __Table`)
	// for rows.Next() {
	// 	var tableName string
	// 	var schema string
	// 	rows.Scan(&tableName, &schema)
	// 	fmt.Printf("Table: %v\n%v\n\n", tableName, schema)
	// }
	// rows.Close()

	rows, err = db.Query(`Select * From nutrients Where nutrient_id == $1`, "255")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var n models.Nutrient
		rows.Scan(&n.FoodID, &n.NutrientID, &n.NutritionValue, &n.Min, &n.Max, &n.DegreesOfFreedom, &n.LowerErrorBound, &n.UpperErrorBound)

		fmt.Printf("ID: %v, %v, %v, %v, %v, %v, %v, %v\n", n.FoodID, n.NutrientID, n.NutritionValue, n.Min, n.Max, n.DegreesOfFreedom, n.LowerErrorBound, n.UpperErrorBound)
	}
	rows.Close()
}
