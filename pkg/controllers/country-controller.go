package controllers

import (
	"net/http"

	"listof/pkg/config"
	"listof/pkg/utils"
)

func GetCountries(w http.ResponseWriter, r *http.Request) {
	countries := config.Countries

	if len(countries) == 0 {
		utils.RespondWithJSON(w, http.StatusNotFound, "No countries found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, countries)
}
