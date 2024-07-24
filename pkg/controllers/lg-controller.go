package controllers

import (
	"listof/pkg/config"
	"listof/pkg/utils"
	"net/http"
)

func GetLGs(w http.ResponseWriter, r *http.Request) {
	
	lgs := config.LocalGovernments
	if len(lgs) == 0 {
		utils.RespondWithJSON(w, http.StatusNotFound, "No local governments found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, lgs)
}
