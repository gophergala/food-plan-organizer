package manage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gophergala/food-plan-organizer/models"
)

type recipeServer struct {
	*sql.DB
}

func (rs *recipeServer) CreateRecipe(rw http.ResponseWriter, req *http.Request) {
	var dec = json.NewDecoder(req.Body)
	var recipe models.Recipe

	if err := dec.Decode(&recipe); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var enc = json.NewEncoder(rw)
	rs.Exec(`INSERT INTO recipes (name, description) VALUES ($1,$2);`, recipe.Name, recipe.Description)
	if err := rs.QueryRow(`SELECT last_insert_rowid() FROM recipes`).Scan(&recipe.ID); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode(map[string]string{"sql_error": fmt.Sprintf("%v", err)})
		return
	}

	enc.Encode(recipe)
}

func (rs *recipeServer) ListRecipes(rw http.ResponseWriter, req *http.Request) {
	var recipes []models.Recipe
	var enc = json.NewEncoder(rw)
	var rows, err = rs.DB.Query(`SELECT * FROM recipes`)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode(map[string]string{"sql_error": fmt.Sprintf("%v", err)})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Recipe
		if err = rows.Scan(&r.ID, &r.Name, &r.Description); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			enc.Encode(map[string]string{"sql_error": fmt.Sprintf("%v", err)})
			return
		}
		recipes = append(recipes, r)
	}

	enc.Encode(recipes)
}

func (rs *recipeServer) DeleteRecipe(rw http.ResponseWriter, req *http.Request) {
	params, _ := url.ParseQuery(req.URL.RawQuery)
	id, ok := params["id"]
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rs.DB.Exec(`DELETE FROM recipes WHERE id = $1`, id[0])
}

func (rs *recipeServer) ShowRecipe(rw http.ResponseWriter, req *http.Request) {
	params, _ := url.ParseQuery(req.URL.RawQuery)
	id, ok := params["id"]
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var r models.Recipe
	var enc = json.NewEncoder(rw)
	if err := rs.DB.QueryRow(`SELECT * FROM recipes WHERE id = $1`, id[0]).Scan(&r.ID, &r.Name, &r.Description); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode(map[string]string{"sql_error": fmt.Sprintf("%v", err)})
		return
	}

	enc.Encode(r)
}

func (rs *recipeServer) UpdateRecipe(rw http.ResponseWriter, req *http.Request) {
	params, _ := url.ParseQuery(req.URL.RawQuery)
	id, ok := params["id"]
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var dec = json.NewDecoder(req.Body)
	var recipe models.Recipe

	if err := dec.Decode(&recipe); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var enc = json.NewEncoder(rw)
	if _, err := rs.Exec(`UPDATE recipes SET name = $1, description = $2 WHERE id = $3;`, recipe.Name, recipe.Description, id[0]); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode(map[string]string{"sql_error": fmt.Sprintf("%v", err)})
		return
	}

	enc.Encode(recipe)
}

func (rs *recipeServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	params, _ := url.ParseQuery(req.URL.RawQuery)

	if req.Method == "GET" {
		if _, ok := params["id"]; ok {
			rs.ShowRecipe(rw, req)
		} else {
			rs.ListRecipes(rw, req)
		}
	} else if req.Method == "POST" {
		rs.CreateRecipe(rw, req)
	} else if req.Method == "PUT" {
		rs.UpdateRecipe(rw, req)
	} else if req.Method == "DELETE" {
		rs.DeleteRecipe(rw, req)
	}
}

func NewRecipeServer(DB *sql.DB) http.Handler {
	return &recipeServer{DB}
}
