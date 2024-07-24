package routes

import (
	"github.com/gorilla/mux"
	"listof/pkg/controllers"
)

func SetupLGRoutes(router *mux.Router) {
	router.HandleFunc("/states/{state_id}/lgs", controllers.GetLGs).Methods("GET")
}
