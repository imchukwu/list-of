package controllers

import (	
    "net/http"
    "strconv"
    "strings"

    "listof/pkg/config"
    "listof/pkg/models"
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

func GetCitiesByLGA(w http.ResponseWriter, r *http.Request) {
    
    pathParts := strings.Split(r.URL.Path, "/")
    if len(pathParts) < 3 || pathParts[1] != "lgs" || pathParts[3] != "cities" {
        utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid URL path")
        return
    }

    lgID, err := strconv.Atoi(pathParts[2])
    if err != nil {
        utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid State ID")
        return
    }

    var cities []models.City
    for _, city := range config.Cities {
        if city.LGID == lgID {
            cities = append(cities, city)
        }
    }

    if len(cities) == 0 {
        utils.RespondWithJSON(w, http.StatusNotFound, "No cities found for the given local government ID")
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, cities)
}
