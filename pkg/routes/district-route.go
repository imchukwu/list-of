package routes

import (
	"github.com/gorilla/mux"
	"listof/pkg/controllers"
)

func SetupDistrictRoutes(router *mux.Router) {
	router.HandleFunc("/cities/{city_id}/districts", controllers.GetDistricts).Methods("GET")
}
