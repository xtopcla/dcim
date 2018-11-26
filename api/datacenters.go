package api

import (
	"fmt"
	"net/http"
)

func (api *API) GetAllDatacenters() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, api.Store.Datacenter().GetAll())
	}
}

func (api *API) InitDatacenters() {
	router := api.Root.PathPrefix("/datacenters").Subrouter()
	router.HandleFunc("", api.GetAllDatacenters()).Methods("GET")
}
