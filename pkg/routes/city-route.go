package routes

import (
	"github.com/gorilla/mux"
	"listof/pkg/controllers"
)

func SetupCityRoutes(router *mux.Router) {
	router.HandleFunc("/cities", controllers.GetCities).Methods("GET")
	router.HandleFunc("/lgs/{lg_id}/cities", controllers.GetCitiesByLGA).Methods("GET")
}
