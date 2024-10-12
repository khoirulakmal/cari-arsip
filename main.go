package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	tmpl, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Printf("Error ! %s", err.Error())
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			fmt.Fprintln(w, "404 Not Found!")
		} else {
			tmpl.Execute(w, nil)
		}
	})
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8000", nil)
}
