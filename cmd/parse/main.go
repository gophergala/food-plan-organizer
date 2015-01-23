package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"sync"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"

	"github.com/cznic/ql"
	"github.com/gophergala/food-plan-organizer/cmd/parse/etl"
	"github.com/gophergala/food-plan-organizer/models"
)
import "flag"

var (
	dataDirectory = flag.String("data.dir", "", "sr27 data directory")
	database      = flag.String("database", "", "path to database file")
)

type Handler func(interface{}, *sql.Tx)

func persistFood(i interface{}, tx *sql.Tx) {
	if fo, ok := i.(etl.Food); ok {
		var dbFood = models.Food{
			ID:                  fo.ID,
			FoodGroupID:         fo.FoodGroupID,
			Name:                fo.Name,
			ShortName:           fo.ShortName,
			CommonName:          fo.CommonName,
			ScientificName:      fo.ScientificName,
			NitrogenFactor:      fo.NitrogenFactor,
			ProteinFactor:       fo.ProteinFactor,
			FatFactor:           fo.FatFactor,
			CarbonhydrateFactor: fo.CarbohydrateFactor,
		}
		if ins := models.InsertFood(&dbFood, tx); ins != nil {
			panic(ins)
		}
	} else {
		panic("expected etl.Food")
	}
}

func persistFoodGroup(i interface{}, tx *sql.Tx) {
	if fg, ok := i.(etl.FoodGroup); ok {
		var dbGroup = models.FoodGroup{
			ID:   fg.GroupID,
			Name: fg.Name,
		}
		if ins := models.InsertFoodGroup(&dbGroup, tx); ins != nil {
			panic(ins)
		}
	} else {
		panic("expected etl.FoodGroup")
	}
}

func persistLangDescription(i interface{}, tx *sql.Tx) {
	if ld, ok := i.(etl.LanguageDescription); ok {
		var dbLangDesc = models.LanguageDescription{
			Code:        ld.FactorCode,
			Description: ld.Description,
		}
		if ins := models.InsertLanguageDescription(&dbLangDesc, tx); ins != nil {
			panic(ins)
		}
	} else {
		panic("expected etl.LangDescription")
	}
}

func persistLang(i interface{}, tx *sql.Tx) {
	if l, ok := i.(etl.Lang); ok {
		var dbL = models.Language{
			NutrientID: l.NutrientID,
			FactorCode: l.FactorCode,
		}
		if ins := models.InsertLanguage(&dbL, tx); ins != nil {
			panic(ins)
		}
	} else {
		panic("expected etl.Lang")
	}
}

func persistWeight(i interface{}, tx *sql.Tx) {
	if w, ok := i.(etl.Weight); ok {
		var dbW = models.Weight{
			NutrientID:  w.NutrientID,
			Seq:         w.Seq,
			Amount:      w.Amount,
			Description: w.Description,
			GramWeight:  w.GramWeight,
		}
		if ins := models.InsertWeight(&dbW, tx); ins != nil {
			panic(ins)
		}
	} else {
		panic("expected etl.Weight")
	}
}

func persistNutrientDefinitions(i interface{}, tx *sql.Tx) {
	if nd, ok := i.(etl.NutrientDefinition); ok {
		var dbND = models.NutrientDefinition{
			NutrientID:    nd.NutrientID,
			Units:         nd.Units,
			Tagname:       nd.Tagname,
			Description:   nd.Description,
			DecimalPlaces: nd.DecimalPlaces,
		}
		if ins := models.InsertNutrientDefinition(&dbND, tx); ins != nil {
			panic(ins)
		}
	} else {
		panic("expected etl.NutrientDefinition")
	}
}

func persistNutrient(i interface{}, tx *sql.Tx) {
	if n, ok := i.(etl.Nutrient); ok {
		var dbN = models.Nutrient{
			FoodID:           n.FoodID,
			NutrientID:       n.NutrientID,
			NutritionValue:   n.NutritionValue,
			Min:              n.Min,
			Max:              n.Max,
			DegreesOfFreedom: n.DegreesOfFreedom,
			LowerErrorBound:  n.LowerErrorBound,
			UpperErrorBound:  n.UpperErrorBound,
		}
		if ins := models.InsertNutrient(&dbN, tx); ins != nil {
			panic(ins)
		}
	} else {
		panic("expected etl.Nutrient")
	}
}

type file struct {
	Name        string
	Encoding    *encoding.Encoding
	Extractor   etl.Extract
	h           Handler
	Description string
}

var files = []file{
	file{
		Name:        "FOOD_DES.txt",
		Description: "Food Description",
		Extractor:   etl.FoodExtractor{},
		h:           persistFood,
	},
	file{
		Name:        "FD_GROUP.txt",
		Description: "Food Group Description",
		Extractor:   etl.FoodGroupExtractor{},
		h:           persistFoodGroup,
	},
	file{
		Name:        "LANGDESC.txt",
		Description: "Langual Factors Description",
		Extractor:   etl.LangDescriptionExtractor{},
		h:           persistLangDescription,
	},
	file{
		Name:        "LANGUAL.txt",
		Description: "Langual Factor",
		Extractor:   etl.LangExtractor{},
		h:           persistLang,
	},
	file{
		Name:        "NUTR_DEF.txt",
		Encoding:    &charmap.ISO8859_3,
		Description: "Nutrient Definition",
		Extractor:   etl.NutrientDefinitionExtractor{},
		h:           persistNutrientDefinitions,
	},
	file{
		Name:        "NUT_DATA.txt",
		Description: "Nutrient Data",
		Extractor:   etl.NutrientExtractor{},
		h:           persistNutrient,
	},
	file{
		Name:        "DERIV_CD.txt",
		Description: "Data Derivation Code Description",
	},
	file{
		Name:        "SRC_CD.txt",
		Description: "Source Code",
	},
	file{
		Name:        "WEIGHT.txt",
		Description: "Weight",
		Extractor:   etl.WeightExtractor{},
		h:           persistWeight,
	},
	file{
		Name:        "DATA_SRC.txt",
		Description: "Sources of Data",
	},
	file{
		Name:        "DATSRCLN.txt",
		Description: "Sources of Data Link",
	},
	file{
		Name:        "FOOTNOTE.txt",
		Description: "Footnote",
	},
}

func runTx(tx *sql.Tx, sql string) {
	if _, err := tx.Exec(sql); err != nil {
		log.Fatalf("Error executing '%v': %v", sql, err)
	}
}

var db *sql.DB

func setupDB(db *sql.DB) {
	var tx, err = db.Begin()
	defer tx.Commit()

	if err != nil {
		log.Fatalf("Failed to open the transaction: %v\n", err)
	}
	for _, sql := range models.CreateFoodGroupTableSQLs {
		runTx(tx, sql)
	}
	for _, sql := range models.CreateFoodTableSQLs {
		runTx(tx, sql)
	}
	for _, sql := range models.CreateLanguageDescriptionTableSQLs {
		runTx(tx, sql)
	}
	for _, sql := range models.CreateLanguageTableSQLs {
		runTx(tx, sql)
	}
	for _, sql := range models.CreateWeightTableSQLs {
		runTx(tx, sql)
	}
	for _, sql := range models.CreateNutrientDefinitionTableSQLs {
		runTx(tx, sql)
	}
	for _, sql := range models.CreateNutrientTableSQLs {
		runTx(tx, sql)
	}
}

func main() {
	flag.Parse()

	if *dataDirectory == "" {
		fmt.Printf("missing data.dir argument!\n")
		os.Exit(1)
	}
	if *database == "" {
		fmt.Printf("missing database argument!\n")
		os.Exit(1)
	}

	if _, err := os.Stat(*database); err == nil {
		os.Remove(*database)
	}

	ql.RegisterDriver()
	var err error
	db, err = sql.Open("ql", *database)
	if err != nil {
		fmt.Printf("Failed opening database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	setupDB(db)

	// health check
	for _, file := range files {
		var f = path.Join(*dataDirectory, "/", file.Name)
		var _, err = os.Stat(f)
		if err != nil {
			fmt.Printf("missing required file: %v\n", f)
			os.Exit(2)
		}
	}

	var wait = sync.WaitGroup{}
	for _, fi := range files {
		if fi.Extractor != nil {
			wait.Add(1)
			go func(fi file) {
				var f, _ = os.Open(path.Join(*dataDirectory, "/", fi.Name))

				var r io.Reader = f
				if fi.Encoding != nil {
					r = transform.NewReader(r, (*fi.Encoding).NewDecoder())
				}

				var ch = make(chan interface{})
				for k := 0; k < runtime.GOMAXPROCS(0)*2; k++ {
					go func(h Handler) {
						var tx, err = db.Begin()
						if err != nil {
							panic(err)
						}
						for i := range ch {
							h(i, tx)
						}
						tx.Commit()
					}(fi.h)
				}
				if err := fi.Extractor.Parse(r, ch); err != nil {
					fmt.Printf("Failed to extract %v: %v\n", fi.Description, err)
				}
				close(ch)
				wait.Done()
			}(fi)
		}
	}
	wait.Wait()
}
