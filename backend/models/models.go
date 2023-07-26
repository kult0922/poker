package models

type Card struct {
	Suit string `json:"suit"`
	Rank int    `json:"rank"`
}

type Hand struct {
	Cards [5]Card `json:"cards"`
	Name  string  `json:"name"`
}
