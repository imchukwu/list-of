package routes

import (
	"github.com/gorilla/mux"
	"listof/pkg/controllers"
)

func SetupCityRoutes(router *mux.Router) {
	router.HandleFunc("/lgs/{lg_id}/cities", controllers.GetCities).Methods("GET")
}
