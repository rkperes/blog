package htmx

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func (h *Handler) SearchPokemon(w http.ResponseWriter, r *http.Request) {
	type PokemonResponse struct {
		Name    string `json:"name"`
		Sprites struct {
			FrontDefault string `json:"front_default"`
		} `json:"sprites"`
	}

	query := strings.ToLower(r.PostFormValue("pokemon"))
	if query == "" {
		_, _ = w.Write([]byte("Please enter a Pokemon name."))
		return
	}

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + query)
	if err != nil || resp.StatusCode != 200 {
		_, _ = w.Write([]byte("Pokemon not found."))
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var pokemon PokemonResponse
	_ = json.Unmarshal(body, &pokemon)

	result := `<img src="` + pokemon.Sprites.FrontDefault + `" alt="` + pokemon.Name + `">`
	_, _ = w.Write([]byte(result))
}
