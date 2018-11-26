package app

import (
	"github.com/gorilla/mux"
	"github.com/xtopcla/dcim/api"
	"github.com/xtopcla/dcim/store"
)

type App struct {
	Store  store.Store
	Router *mux.Router
}

func (app *App) InitApi(router *mux.Router, store store.Store) *api.API {
	api := &api.API{
		Root:  router,
		Store: store,
	}
	api.Init()
	return api
}

func NewApp(router *mux.Router, store store.Store) *App {
	app := &App{
		Router: router,
		Store:  store,
	}
	app.Router = router
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	app.InitApi(apiRouter, store)
	return app
}
