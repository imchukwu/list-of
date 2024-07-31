package controllers

import (	
    "net/http"
    "strconv"
    "strings"

    "listof/pkg/config"
    "listof/pkg/models"
    "listof/pkg/utils"
)

func GetLGs(w http.ResponseWriter, r *http.Request) {
	
	lgs := config.LocalGovernments
	if len(lgs) == 0 {
		utils.RespondWithJSON(w, http.StatusNotFound, "No local governments found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, lgs)
}

func GetLGsByState(w http.ResponseWriter, r *http.Request) {
    
    pathParts := strings.Split(r.URL.Path, "/")
    if len(pathParts) < 3 || pathParts[1] != "states" || pathParts[3] != "lgs" {
        utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid URL path")
        return
    }

    stateID, err := strconv.Atoi(pathParts[2])
    if err != nil {
        utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid State ID")
        return
    }

    var lgs []models.LocalGovernment
    for _, lg := range config.LocalGovernments {
        if lg.StateID == stateID {
            lgs = append(lgs, lg)
        }
    }

    if len(lgs) == 0 {
        utils.RespondWithJSON(w, http.StatusNotFound, "No local governments found for the given state ID")
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, lgs)
}
