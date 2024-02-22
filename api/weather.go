package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler for fetching weather data
func (a *API) GetWeather() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read the lat and lon from the URL
		vars := mux.Vars(r)
		lat := vars["lat"]
		lon := vars["lon"]

		// get the weather data
		res, err := a.clients.GetWeather(lat, lon)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		weatherStr := fmt.Sprintf("The weather in %v, %v is %v with a temperature of %v", res.Name, res.Sys.Country, res.Weather[0].Description, res.Main.Temp)

		// write the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(weatherStr)
	}
}
