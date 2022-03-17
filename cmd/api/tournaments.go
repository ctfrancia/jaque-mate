package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createTournamentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new tournament")
}

func (app *application) showTournamentHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
