package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/cznic/ql"
	"github.com/gophergala/food-plan-organizer/cmd/serve/search"
)

var (
	listen   = flag.String("listen", ":8080", "Port to listen on")
	database = flag.String("database", "default.db", "QL Database name")
)

func newServer(db *sql.DB) http.Handler {
	var r = http.NewServeMux()
	r.Handle("/search/food/", search.NewSearchServer(db))
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

	ql.RegisterDriver()
	var db, err = sql.Open("ql", *database)
	if err != nil {
		log.Fatalf("Failed to open database: %v\n", err)
	}
	defer db.Close()

	log.Printf("Listening on %v", *listen)
	http.ListenAndServe(*listen, newServer(db))
}
