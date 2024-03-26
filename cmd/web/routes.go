package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/event/view", app.eventView)
	mux.HandleFunc("/event/create", app.eventCreate)
	mux.HandleFunc("/shop/view", app.shopView)
	return mux
}
