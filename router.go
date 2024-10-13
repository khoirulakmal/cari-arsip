package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("ui/assets/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/cari", app.cari)

	return mux
}
