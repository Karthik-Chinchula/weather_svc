package clients

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"weather_svc/models"
)

const (
	weatherURL = "%v/data/2.5/weather?lat=%v&lon=%v&appid=%v"
)

// Clients is a struct that holds the config
type Clients struct {
	config models.Config
}

// NewClients will create a new instance of Clients
func NewClients(config models.Config) *Clients {
	return &Clients{config: config}
}

func (c *Clients) GetWeather(lat, lon string) (models.Response, error) {
	// create a new instance of the response struct
	var res models.Response

	// build the URL
	url := fmt.Sprintf(weatherURL, c.config.APIURL, lat, lon, c.config.ApiKey)

	log.Printf("Client URL: %v", url)

	// make the request
	response, err := http.Get(url)
	if err != nil {
		return res, err
	}
	defer response.Body.Close()

	// decode the response
	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		return res, err
	}

	return res, nil
}
