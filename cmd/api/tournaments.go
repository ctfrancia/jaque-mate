package main

import (
	"fmt"
	"net/http"
	"time"

	"jaque-mate.ctfrancia.es/internal/data"
	"jaque-mate.ctfrancia.es/internal/validator"
)

func (app *application) createTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TournamentName string   `json:"tournamentName"`
		Tags           []string `json:"tags"`
		// LastDateOfRegistration time.Time
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	tournament := &data.Tournament{
		TournamentName: input.TournamentName,
		Tags:           input.Tags,
	}

	v := validator.New()

	if data.ValidateTournament(v, tournament); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
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
