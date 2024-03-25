package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HomePage")
}

func (app *application) eventView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ViewPage")
}

func (app *application) eventCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CreatePage")
}
