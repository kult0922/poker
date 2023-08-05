package deal

import (
	"math/rand"
	"time"

	"github.com/kult0922/poker/backend/domain/models"
)

func CommunityCard() [5]models.Card {

	var communityCard []models.Card

	for i := 0; i < 5; i++ {
		communityCard = append(communityCard, Draw(communityCard))
	}

	return [5]models.Card(communityCard)
}

func Draw(exclusionCards []models.Card) models.Card {
	stock := []models.Card{
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

	for _, exclusionCard := range exclusionCards {
		for i, card := range stock {
			if card == exclusionCard {
				stock = append(stock[:i], stock[i+1:]...)
			}
		}
	}

	seed := time.Now().UnixNano()
	rand.Seed(seed)
	randomIndex := rand.Int31n(int32(len(stock)))

	return stock[randomIndex]
}
