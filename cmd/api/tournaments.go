package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"jaque-mate.ctfrancia.es/internal/data"
)

func (app *application) createTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TournamentName string   `json:"tournamentName"`
		Tags           []string `json:"tags"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showTournamentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundRequest(w, r)
		return
	}

	tournament := data.Tournament{
		ID:             id,
		CreatedAt:      time.Now(),
		TournamentName: "Cool Tournament Name",
		Tags:           []string{"fast", "barcelona", "rated"},
		Version:        1,
		Empty:          "foo",
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"tournament": tournament}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
