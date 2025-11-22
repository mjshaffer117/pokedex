// pokedex
package main

type Config struct {
    NextUrl string
    PrevUrl string
}

func main() {
    appConfig := Config{
        NextUrl: "https://pokeapi.co/api/v2/location-area/",
        PrevUrl: "",
    }
    startRepl(&appConfig)
}
