package models

type Card struct {
	Suit string `json:"suit"`
	Rank int    `json:"rank"`
}

type Hand struct {
	Cards [5]Card `json:"cards"`
	Name  string  `json:"name"`
	Rank  int     `json:"rank"`
}

func GetStock() []Card {
	return []Card{
		{Suit: "spade", Rank: 1},
		{Suit: "spade", Rank: 2},
		{Suit: "spade", Rank: 3},
		{Suit: "spade", Rank: 4},
		{Suit: "spade", Rank: 5},
		{Suit: "spade", Rank: 6},
		{Suit: "spade", Rank: 7},
		{Suit: "spade", Rank: 8},
		{Suit: "spade", Rank: 9},
		{Suit: "spade", Rank: 10},
		{Suit: "spade", Rank: 11},
		{Suit: "spade", Rank: 12},
		{Suit: "spade", Rank: 13},
		{Suit: "heart", Rank: 1},
		{Suit: "heart", Rank: 2},
		{Suit: "heart", Rank: 3},
		{Suit: "heart", Rank: 4},
		{Suit: "heart", Rank: 5},
		{Suit: "heart", Rank: 6},
		{Suit: "heart", Rank: 7},
		{Suit: "heart", Rank: 8},
		{Suit: "heart", Rank: 9},
		{Suit: "heart", Rank: 10},
		{Suit: "heart", Rank: 11},
		{Suit: "heart", Rank: 12},
		{Suit: "heart", Rank: 13},
		{Suit: "diamond", Rank: 1},
		{Suit: "diamond", Rank: 2},
		{Suit: "diamond", Rank: 3},
		{Suit: "diamond", Rank: 4},
		{Suit: "diamond", Rank: 5},
		{Suit: "diamond", Rank: 6},
		{Suit: "diamond", Rank: 7},
		{Suit: "diamond", Rank: 8},
		{Suit: "diamond", Rank: 9},
		{Suit: "diamond", Rank: 10},
		{Suit: "diamond", Rank: 11},
		{Suit: "diamond", Rank: 12},
		{Suit: "diamond", Rank: 13},
		{Suit: "club", Rank: 1},
		{Suit: "club", Rank: 2},
		{Suit: "club", Rank: 3},
		{Suit: "club", Rank: 4},
		{Suit: "club", Rank: 5},
		{Suit: "club", Rank: 6},
		{Suit: "club", Rank: 7},
		{Suit: "club", Rank: 8},
		{Suit: "club", Rank: 9},
		{Suit: "club", Rank: 10},
		{Suit: "club", Rank: 11},
		{Suit: "club", Rank: 12},
		{Suit: "club", Rank: 13},
	}
}
