package main

import (
	"fmt"
	"html/template"
	"path/filepath"
)

type templateData struct {
	Page    string
	Counter int
}

// Parse template and save it into a memory
func parseTemplate() (map[string]*template.Template, error) {
	templateCache := make(map[string]*template.Template)
	pages, err := filepath.Glob("./ui/html/page/*.html")
	if err != nil {
		fmt.Printf("Error! %s", err.Error())
		return nil, err
	}
	files, err := filepath.Glob("./ui/html/*.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		// Create a slice containing the filepaths for our base template, any
		// partials and the page.
		parseFile := append(files, page)
		parseFile = append(parseFile, "./ui/html/page/cari.html")
		ts, err := template.New(name).ParseFiles(parseFile...)
		if err != nil {
			return nil, err
		}
		templateCache[name] = ts

	}
	fmt.Println(templateCache)
	return templateCache, nil
}
