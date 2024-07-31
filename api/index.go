// 

package main

import (
	"net/http"

	"listof/pkg/routes"
	"listof/pkg/config"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := config.SetupApp()
	routes.SetupRoutes(router)
	router.ServeHTTP(w, r)
}


