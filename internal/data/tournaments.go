package data

import (
	"time"
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
