package controllers

import (
	"net/http"
	"listof/pkg/config"
	"listof/pkg/utils"
)

func GetStates(w http.ResponseWriter, r *http.Request) {
	
	states := config.States
	if len(states) == 0 {
		utils.RespondWithJSON(w, http.StatusNotFound, "No states found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, states)
}
