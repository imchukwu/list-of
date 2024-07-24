package config

import (
	"log"

	"listof/pkg/models"

	"github.com/gorilla/mux"
)

var (
	Countries       []models.Country
	States          []models.State
	LocalGovernments []models.LocalGovernment
	Cities          []models.City
	Districts       []models.District
)

func SetupApp() *mux.Router {
	router := mux.NewRouter()
	initData()
	return router
}

func initData() {
	log.Println("Initializing static data...")

	Countries = []models.Country{
		{ID: 1, Name: "Nigeria"},
	}

	States = []models.State{
		{ID: 1, Name: "Abia", CountryID: 1},
		{ID: 2, Name: "Adamawa", CountryID: 1},
		{ID: 3, Name: "Akwa Ibom", CountryID: 1},			
		{ID: 4, Name: "Anambra", CountryID: 1},
		{ID: 5, Name: "Bauchi", CountryID: 1},
		{ID: 6, Name: "Bayelsa", CountryID: 1},				
		{ID: 7, Name: "Benue", CountryID: 1},
		{ID: 8, Name: "Borno", CountryID: 1},
		{ID: 9, Name: "Cross River", CountryID: 1},			
		{ID: 10, Name: "Delta", CountryID: 1},
		{ID: 11, Name: "Ebonyi", CountryID: 1},
		{ID: 12, Name: "Edo", CountryID: 1},				
		{ID: 13, Name: "Ekiti", CountryID: 1},
		{ID: 14, Name: "Enugu", CountryID: 1},
		{ID: 15, Name: "Gombe", CountryID: 1},				
		{ID: 16, Name: "Imo", CountryID: 1},
		{ID: 17, Name: "Jigawa", CountryID: 1},
		{ID: 18, Name: "Kaduna", CountryID: 1},				
		{ID: 19, Name: "Kano", CountryID: 1},
		{ID: 20, Name: "Katsina", CountryID: 1},
		{ID: 21, Name: "Kebbi", CountryID: 1},				
		{ID: 22, Name: "Kogi", CountryID: 1},
		{ID: 23, Name: "Kwara", CountryID: 1},
		{ID: 24, Name: "Lagos", CountryID: 1},				
		{ID: 25, Name: "Nasarawa", CountryID: 1},
		{ID: 26, Name: "Niger", CountryID: 1},
		{ID: 27, Name: "Ogun", CountryID: 1},				
		{ID: 28, Name: "Ondo", CountryID: 1},
		{ID: 29, Name: "Osun", CountryID: 1},
		{ID: 30, Name: "Oyo", CountryID: 1},				
		{ID: 31, Name: "Plateau", CountryID: 1},
		{ID: 32, Name: "Rivers", CountryID: 1},
		{ID: 33, Name: "Sokoto", CountryID: 1},				
		{ID: 34, Name: "Taraba", CountryID: 1},
		{ID: 35, Name: "Yobe", CountryID: 1},
		{ID: 36, Name: "Zamfara", CountryID: 1},			
		{ID: 37, Name: "Abuja", CountryID: 1},
	}

	LocalGovernments = []models.LocalGovernment{
		{ID: 1, Name: "Lagos Mainland", StateID: 1},
		{ID: 2, Name: "Lagos Island", StateID: 1},
		{ID: 3, Name: "Fagge", StateID: 2},
		{ID: 4, Name: "Gwale", StateID: 2},
		{ID: 5, Name: "Port Harcourt", StateID: 3},
		{ID: 6, Name: "Obio-Akpor", StateID: 3},
	}

	Cities = []models.City{
		{ID: 1, Name: "Lagos City", LGID: 1},
		{ID: 2, Name: "Kano City", LGID: 3},
		{ID: 3, Name: "Port Harcourt City", LGID: 5},
	}

	Districts = []models.District{
		{ID: 1, Name: "Ikeja", CityID: 1},
		{ID: 2, Name: "Victoria Island", CityID: 1},
		{ID: 3, Name: "Sabon Gari", CityID: 2},
		{ID: 4, Name: "GRA Phase 2", CityID: 3},
	}
}
