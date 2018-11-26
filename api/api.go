package api

import (
	"github.com/gorilla/mux"
	"github.com/xtopcla/dcim/store"
)

type API struct {
	Root  *mux.Router
	Store store.Store
}

func (api *API) Init() {
	api.InitDatacenters()
}
