package main

import (
	"fmt"
	"net/http"
	"time"

	"jaque-mate.ctfrancia.es/internal/data"
)

func (app *application) createTournamentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new tournament")
}

func (app *application) showTournamentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
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
		app.logger.Println(err)
		http.Error(w, "Server encountered an issue", http.StatusInternalServerError)
	}
}
