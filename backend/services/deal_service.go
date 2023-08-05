package services

import (
	"github.com/kult0922/poker/backend/domain/deal"
	"github.com/kult0922/poker/backend/domain/models"
)

func (s *MyAppService) GetCommunityCardService() [5]models.Card {

	return deal.CommunityCard()
}
