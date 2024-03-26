package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Title       string    `json:"title"`
	Category    []string  `json:"category"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	event := Event{}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	file := "./ui/html/home.html"

	tmpl, err := template.ParseFiles(file)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error Here", 500)
		return
	}
	log.Print("Parsed file")

	if err := tmpl.Execute(w, event); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) eventView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Events View Page coming soon!")
}

func (app *application) eventCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Events Create Page coming soon!")
}

func (app *application) shopView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Shop Page coming soon!")
}
