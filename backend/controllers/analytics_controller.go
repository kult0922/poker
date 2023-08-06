package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kult0922/poker/backend/controllers/services"
)

type AnalyticsController struct {
	service services.AnalyticsServicer
}

func NewAnalyticsController(s services.AnalyticsServicer) *AnalyticsController {
	return &AnalyticsController{service: s}
}

func (c *AnalyticsController) PreflopHandRangeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("PreflopHandRangeHandler")

	preflopHandRange := c.service.GetPreflopHandRange()

	json.NewEncoder(w).Encode(preflopHandRange)
}
