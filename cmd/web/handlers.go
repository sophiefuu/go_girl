package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
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

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) eventView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Events View Page coming soon!")
}

func (app *application) eventCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Events Create Page coming soon!")
}

func (app *application) sportsView(w http.ResponseWriter, r *http.Request) {
	file := "./ui/html/halfMarathonTraining.html"

	tmpl, err := template.ParseFiles(file)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error Here", 500)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) shopView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Shop Page coming soon!")
}
