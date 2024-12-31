package handlers

import (
	"html/template"
	"net/http"
	"pencarian_elektronik/services"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	products, err := services.SearchProducts(query)
	if err != nil {
		if err == services.ErrQueryTooShort {
			http.Error(w, "Query too short. Please provide at least 3 characters.", http.StatusBadRequest)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	err = tmpl.Execute(w, products)
	if err != nil {
		http.Error(w, "Error rendering search results", http.StatusInternalServerError)
	}
}

func SearchPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/search.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering search page", http.StatusInternalServerError)
	}
}
