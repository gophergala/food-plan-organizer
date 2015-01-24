package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"time"

	_ "github.com/gophergala/food-plan-organizer/_third_party/github.com/mattn/go-sqlite3"
	"github.com/gophergala/food-plan-organizer/cmd/serve/search"
	"github.com/gophergala/food-plan-organizer/cmd/serve/show"
)

var (
	listen   = flag.String("listen", ":8080", "Port to listen on")
	database = flag.String("database", "default.db", "QL Database name")
)

func newServer(db *sql.DB) http.Handler {
	var r = http.NewServeMux()
	r.Handle("/search/food/", search.NewFoodSearchServer(db))
	r.Handle("/food/", show.NewFoodShowServer(db))
	return logHandler(jsonHandler(r))
}

func logHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}

func jsonHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	}
}

func main() {
	flag.Parse()

	var db, err = sql.Open("sqlite3", *database)
	if err != nil {
		log.Fatalf("Failed to open database: %v\n", err)
	}
	defer db.Close()

	log.Printf("Listening on %v", *listen)
	http.ListenAndServe(*listen, newServer(db))
}
