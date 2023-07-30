package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kult0922/go-react-blog/backend/controllers/services"
)

type DealController struct {
	service services.DealServicer
}

func NewDealController(s services.DealServicer) *DealController {
	return &DealController{service: s}
}

func (c *DealController) CommunityCardHandler(w http.ResponseWriter, req *http.Request) {

	communityCard := c.service.GetCommunityCardService()

	json.NewEncoder(w).Encode(communityCard)
}
