package controllers

import (
	"net/http"

	"listof/pkg/config"
	"listof/pkg/utils"
)

func GetDistricts(w http.ResponseWriter, r *http.Request) {
	
	districts := config.Districts
	if len(districts) == 0 {
		utils.RespondWithJSON(w, http.StatusNotFound, "No districts found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, districts)
}