package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kult0922/poker/backend/api/middlewares"
	"github.com/kult0922/poker/backend/controllers"
	"github.com/kult0922/poker/backend/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	HandCon := controllers.NewHandController(ser)
	DealCon := controllers.NewDealController(ser)
	AnalyticsCon := controllers.NewAnalyticsController(ser)

	r := mux.NewRouter()
	r.Use(middlewares.Cors)
	r.HandleFunc("/hand", HandCon.HandHandler).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/community-card", DealCon.CommunityCardHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/analytics/preflop", AnalyticsCon.PreflopHandRangeHandler).Methods(http.MethodGet, http.MethodOptions)

	r.Use(middlewares.LoggingMiddleware)
	return r
}
