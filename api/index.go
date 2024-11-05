package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

// Pokemon represents the structure of each Pokemon entry.
type Pokemon struct {
	ImageBlob string `json:"imageBlob"`
	Name      string `json:"name"`
	ID        int    `json:"pokedexId"`
}

// api fetches data from the API and returns a list of Pokemon structs.
func api() ([]Pokemon, error) {
	url := "https://node-server-seven-chi.vercel.app/pokemon"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var pokemons []Pokemon
	if err := json.Unmarshal(body, &pokemons); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return pokemons, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	pokemons, err := api()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("error parsing template: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, pokemons); err != nil {
		http.Error(w, fmt.Sprintf("error executing template: %v", err), http.StatusInternalServerError)
	}
}

// Exported function for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
