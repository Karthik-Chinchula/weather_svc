package main

import (
	"log"
	"net/http"
	"weather_svc/api"
	"weather_svc/clients"
	"weather_svc/config"
)

func main() {
	// Read the config file
	cfg, err := config.ReadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new instance of the clients
	cli := clients.NewClients(cfg)

	// Create a new instance of the API
	api := api.NewAPI(cli)

	// Log that the server has started
	log.Println("Server started on port 8080")

	// Start the server
	http.ListenAndServe(":8080", api.GetRouter())
}
