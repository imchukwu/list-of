package routes

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	SetupCountryRoutes(router)
	SetupStateRoutes(router)
	SetupLGRoutes(router)
	SetupCityRoutes(router)
	SetupDistrictRoutes(router)
}
