package api

import (
	"net/http"
	"weather_svc/clients"

	"github.com/gorilla/mux"
)

type API struct {
	clients *clients.Clients
	router  *mux.Router
}

func NewAPI(clients *clients.Clients) *API {
	api := &API{
		clients: clients,
	}

	api.router = mux.NewRouter()
	api.router.HandleFunc("/weather/{lat}/{lon}", api.GetWeather()).Methods(http.MethodGet)

	return api
}

func (a *API) GetRouter() *mux.Router {
	return a.router
}
