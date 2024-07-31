package routes

import (
	"github.com/gorilla/mux"
	"listof/pkg/controllers"
)

func SetupStateRoutes(router *mux.Router) {
	router.HandleFunc("/states", controllers.GetStates).Methods("GET")
	router.HandleFunc("/countries/{country_id}/states", controllers.GetStatesByCountry).Methods("GET")
}
