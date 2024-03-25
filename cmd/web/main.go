package main

import (
	"flag"
	"log"
	"net/http"
)

type application struct {
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	app := &application{}
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)

}
