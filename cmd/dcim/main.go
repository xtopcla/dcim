package main

import (
	"github.com/gorilla/mux"
	"github.com/xtopcla/dcim/app"
	"github.com/xtopcla/dcim/store/sqlstore"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	app.NewApp(router, &sqlstore.SqlStore{})
	log.Fatal(http.ListenAndServe(":8000", router))
}
