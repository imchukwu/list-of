// cmd/main/main.go

package main

import (
	"log"
	"net/http"

	"listof/pkg/routes"
	"listof/pkg/config"
)

func main() {
	router := config.SetupApp()
	
	// Initialize routes
	routes.SetupRoutes(router)

	// Start server
	log.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal(err)
	}
}
