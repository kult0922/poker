package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kult0922/poker/backend/apperrors"
	"github.com/kult0922/poker/backend/controllers/services"
	"github.com/kult0922/poker/backend/domain/models"
)

type HandController struct {
	service services.HandServicer
}

func NewHandController(s services.HandServicer) *HandController {
	return &HandController{service: s}
}

func (c *HandController) HandHandler(w http.ResponseWriter, req *http.Request) {
	var requestHand [5]models.Card

	if err := json.NewDecoder(req.Body).Decode(&requestHand); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	hand := c.service.GetHandService(requestHand)

	json.NewEncoder(w).Encode(hand)
}
