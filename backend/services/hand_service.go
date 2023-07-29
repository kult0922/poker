package services

import (
	"github.com/kult0922/go-react-blog/backend/models"
)

func (s *MyAppService) GetHandService(cards [5]models.Card) models.Hand {

	return models.Hand{
		Cards: cards,
		Name:  models.HandName(cards),
	}
}
