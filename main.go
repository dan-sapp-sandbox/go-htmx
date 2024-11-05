package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	pokemons, err := Api()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	queryIds := r.URL.Query().Get("id")
	filteredPokemons := []Pokemon{}

	if queryIds != "" {
		idList := strings.Split(queryIds, ",")
		for _, p := range pokemons {
			for _, id := range idList {
				if strconv.Itoa(p.ID) == id {
					filteredPokemons = append(filteredPokemons, p)
					break
				}
			}
		}
	} else {
		// If no query parameter, you could return an empty list or all pokemons
		filteredPokemons = pokemons // Replace with an empty list if needed
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("error parsing template: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, pokemons); err != nil {
		http.Error(w, fmt.Sprintf("error executing template: %v", err), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
