package services

import (
	"github.com/kult0922/poker/backend/domain/hand"
	"github.com/kult0922/poker/backend/domain/models"
)

func (s *MyAppService) GetHandService(cards [5]models.Card) models.Hand {
	handName := hand.HandName(cards)

	return models.Hand{
		Cards: cards,
		Name:  handName,
		Rank:  hand.HandRank(s.rankMap, s.rankMapFlush, cards, handName),
	}
}
