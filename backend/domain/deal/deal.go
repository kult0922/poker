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
	stock := models.GetStock()

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
