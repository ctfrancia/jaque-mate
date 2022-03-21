package data

import (
	"time"

	"jaque-mate.ctfrancia.es/internal/validator"
)

// Tournament struct defines the object structure of each individual Tournament
type Tournament struct {
	ID             int64     `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	TournamentName string    `json:"tournamentName"`
	Tags           []string  `json:"tags,omitempty"`
	Version        int32     `json:"version"`
	Empty          string    `json:"-"` // "-" always omits this struct field
}

// ValidateTournament will validate the input of the tournament desired to be saved
func ValidateTournament(v *validator.Validator, tournament *Tournament) {
	v.Check(tournament.TournamentName != "", "tournament name", "must be provided")
	v.Check(len(tournament.TournamentName) <= 500, "tournament name", "must be more than 500 bits long")

	v.Check(tournament.Tags != nil, "tags", "must be provided")
	v.Check(validator.Unique(tournament.Tags), "tags", "must not contain duplicate values")
}
