package services

import (
	"github.com/kult0922/go-react-blog/backend/models"
)

func (s *MyAppService) GetCommunityCardService() [5]models.Card {

	return models.CommunityCard()
}
