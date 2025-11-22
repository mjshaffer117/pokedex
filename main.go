// pokedex
// main entry code

package main

import (
    "time"
    "github.com/mjshaffer117/pokedex/internal/pokeapi"
)


func main() {
    pokeClient := pokeapi.NewClient(5 * time.Second)
    cfg := &config {
        pokeapiClient: pokeClient,
    }

    startRepl(cfg)
}
