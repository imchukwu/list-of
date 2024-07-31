package controllers

import (
    "net/http"
    "strconv"
    "strings"

    "listof/pkg/config"
    "listof/pkg/models"
    "listof/pkg/utils"
)


func GetStates(w http.ResponseWriter, r *http.Request) {
	
	states := config.LocalGovernments
	if len(states) == 0 {
		utils.RespondWithJSON(w, http.StatusNotFound, "No states found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, states)
}

func GetStatesByCountry(w http.ResponseWriter, r *http.Request) {
    
    pathParts := strings.Split(r.URL.Path, "/")
    if len(pathParts) < 3 || pathParts[1] != "countries" || pathParts[3] != "states" {
        utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid URL path")
        return
    }

    countryID, err := strconv.Atoi(pathParts[2])
    if err != nil {
        utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid Country ID")
        return
    }

    var states []models.State
    for _, state := range config.States {
        if state.CountryID == countryID {
            states = append(states, state)
        }
    }

    if len(states) == 0 {
        utils.RespondWithJSON(w, http.StatusNotFound, "No states found for the given country ID")
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, states)
}
