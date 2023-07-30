package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kult0922/go-react-blog/backend/api/middlewares"
	"github.com/kult0922/go-react-blog/backend/controllers"
	"github.com/kult0922/go-react-blog/backend/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	HandCon := controllers.NewHandController(ser)
	DealCon := controllers.NewDealController(ser)

	r := mux.NewRouter()
	r.Use(middlewares.Cors)
	r.HandleFunc("/hand", HandCon.HandHandler).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/community-card", DealCon.CommunityCardHandler).Methods(http.MethodGet, http.MethodOptions)

	r.Use(middlewares.LoggingMiddleware)
	return r
}
