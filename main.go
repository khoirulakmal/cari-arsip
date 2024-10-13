package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
}

func main() {

	errorLog := log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

	// Parse template cache
	tmpl, err := parseTemplate()
	if err != nil {
		errorLog.Printf(err.Error())
	} else {
		infoLog.Printf("Parsing template success!")
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: tmpl,
	}

	srvr := &http.Server{
		Addr:    ":8000",
		Handler: app.routes(),
	}
	err = srvr.ListenAndServe()
	if err != nil {
		errorLog.Printf(err.Error())
	}
}
