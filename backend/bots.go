package main

// Bots struct for the bots in the robo-wars
type Bots struct {
	Name        string     `json:"name"`
	Team        string     `json:"name"`
	Category    Categories `json:"category"`
	Description string     `json:"description"`
	Youtube     string     `json:"video"`
	Image       []byte     `json:"img"`
}

// Categories define the type of Categories
type Categories int

// Enum for Different Categories
const (
	Fifteen Categories = 15
	Thirty             = 30
	Sixty              = 60
)
