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
	Con := controllers.NewHandController(ser)

	r := mux.NewRouter()
	r.Use(middlewares.Cors)
	r.HandleFunc("/hand", Con.HandHandler).Methods(http.MethodPost, http.MethodOptions)

	r.Use(middlewares.LoggingMiddleware)
	return r
}
