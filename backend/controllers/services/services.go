package services

import "github.com/kult0922/poker/backend/domain/models"

type HandServicer interface {
	GetHandService(cards [5]models.Card) models.Hand
}

type DealServicer interface {
	GetCommunityCardService() [5]models.Card
}

type AnalyticsServicer interface {
	GetPreflopHandRange() map[string]int
}
