package controllers

import (
	"net/http"

	"listof/pkg/config"
	"listof/pkg/utils"
)


func GetCities(w http.ResponseWriter, r *http.Request) {
	
	cities := config.Cities
	if len(cities) == 0 {
		utils.RespondWithJSON(w, http.StatusNotFound, "No cities found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, cities)
}