package main

import (
	"log"
	"net/http"

	"github.com/sophiefuu/go_getter/internal/models"
)

type application struct {
	eventlist *models.EventlistModel
}

func main() {
	addr := "8080" //flag.String("addr", ":8080", "HTTP network address")
	app := &application{}
	srv := &http.Server{
		Addr:    ":" + addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)

}
