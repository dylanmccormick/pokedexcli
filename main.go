package main

import (
	"time"

	"github.com/dylanmccormick/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       map[string]pokeapi.Pokemon{},
	}
	StartRepl(cfg)

}
