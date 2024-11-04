package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Pokemon represents the structure of each Pokemon entry.
type Pokemon struct {
	SpriteUrl string `json:"spriteUrl"` // URL for the image
	Name      string `json:"name"`      // Alt text or name for the image
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
