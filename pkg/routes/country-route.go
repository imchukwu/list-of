package routes

import (
	"github.com/gorilla/mux"
	"listof/pkg/controllers"
)

func SetupCountryRoutes(router *mux.Router) {
	router.HandleFunc("/countries", controllers.GetCountries).Methods("GET")
}

