package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "./templates/test.page.gohtml")
	})

	// need to be 80 on production!!!!
	port := ":8000"
	fmt.Printf("Starting front end service on port %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {

	templateSlice := []string{
		t,
		"./templates/base.layout.gohtml",
		"./templates/header.partial.gohtml",
		"./templates/footer.partial.gohtml",
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
