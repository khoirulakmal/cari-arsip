package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// The serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, e error) {
	trace := fmt.Sprintf("%s\n%s", e.Error(), debug.Stack())
	app.errorLog.Printf(trace)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// Render page based on data requested
func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	pageHTML := fmt.Sprintf("%s.html", page)
	ts, ok := app.templateCache[pageHTML]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}
	w.WriteHeader(status)
	err := ts.ExecuteTemplate(w, data.Page, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) generateTemplateData() *templateData {
	return &templateData{
		Page: "base",
	}
}
