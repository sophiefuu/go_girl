package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/sophiefuu/go_girl/internal/models"
)

type application struct {
	eventlist *models.EventlistModel
}

func main() {
	addr := flag.String("addr", "8080", "HTTP network address") //os.Getenv("PORT")
	app := &application{}
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)

}
