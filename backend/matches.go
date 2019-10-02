package main

// MatchStatus return different codes for different match status
type MatchStatus int

const (
	Upcoming MatchStatus = 0 + iota
	Red
	Blue
	Draw
)

// Matches struct for matches
type Matches struct {
	MatchID  string      `json:"id"`
	BlueTeam string      `json:"blue"`
	RedTeam  string      `json:"red"`
	Category Categories  `json:"category"`
	Status   MatchStatus `json:"status"`
}
