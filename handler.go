package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	data := app.generateTemplateData()
	app.render(w, http.StatusOK, "main", data)
}

func (app *application) cari(w http.ResponseWriter, r *http.Request) {
	data := app.generateTemplateData()
	data.Page = "cari"
	app.render(w, http.StatusOK, "main", data)
}
