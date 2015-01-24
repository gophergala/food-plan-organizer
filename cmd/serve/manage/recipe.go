package manage

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
)

type recipeServer struct {
	*sql.DB
}

func (rs *recipeServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	params, _ := url.ParseQuery(req.URL.RawQuery)

	if req.Method == "GET" {
		if _, ok := params["id"]; ok {
			fmt.Printf("SHOW\n")
		} else {
			fmt.Printf("INDEX\n")
		}
		// GET  -> INDEX
		// GET ?id= -> SHOW
	} else if req.Method == "POST" {
		// POST / -> CREATE

	} else if req.Method == "PUT" {
		// PUT ?id= -> UPDATE

	} else if req.Method == "DELETE" {
		// DELETE ?id= DESTROY

	}
}

func NewRecipeServer(DB *sql.DB) http.Handler {
	return &recipeServer{DB}
}
