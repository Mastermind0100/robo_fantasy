package main

// User struct for the user
type User struct {
	FirstName   string           `json:"firstName"`
	LastName    string           `json:"lastName"`
	Email       string           `json:"email"`
	Predictions []UserPrediction `json:"pred"`
}

type UserPrediction struct {
	MatchID string `json:"matchID"`
	Team    string `json:"team"`
	Result  bool   `json:"result"`
}
