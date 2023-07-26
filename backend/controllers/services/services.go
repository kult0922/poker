package services

import "github.com/kult0922/go-react-blog/backend/models"

type HandServicer interface {
	GetHandService(cards [5]models.Card) models.Hand
}
