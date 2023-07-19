package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	fmt.Println("Starting front end service on port 3001")
	err := http.ListenAndServe("localhost:3001", nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, templateName string) {
	templateBasePath := "./cmd/web/templates/"
	partials := []string{
		templateBasePath + "base.layout.gohtml",
		templateBasePath + "header.partial.gohtml",
		templateBasePath + "footer.partial.gohtml",
	}

	templateSlices := []string{templateBasePath + templateName}
	templateSlices = append(templateSlices, partials...)

	tmpl, err := template.ParseFiles(templateSlices...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
