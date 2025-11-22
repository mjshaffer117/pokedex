// pokedex
// List of valid functions within the pokedex cli tool

package main

import (
    "fmt"
    "os"
    "pokedex/internal/pokeapi"
)


func commandExit(cfg *Config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}


func commandHelp(cfg *Config) error {
    fmt.Println()
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, commandInfo := range getCommands() {
        fmt.Printf("%s: %s\n", commandInfo.name, commandInfo.description)
    }
    fmt.Println()
    return nil
}


func commandMap(cfg *Config) error {
    if cfg.NextUrl == "" {
        fmt.Println("Last page reached, use 'mapb' to back-track")
        return nil
    }

    locationsData, err := pokeapi.GetLocations(cfg.NextUrl)
    if err != nil {
        return fmt.Errorf("Error returning location data.")
    }
    for _, location := range locationsData.Results {
        fmt.Println(location.Name)
    }

    cfg.NextUrl = locationsData.Next
    cfg.PrevUrl = locationsData.Prev

    return nil
}


func commandMapB(cfg *Config) error {
    if cfg.PrevUrl == "" {
        fmt.Println("First page reached, use 'map' to jump to the next page")
        return nil
    }

    locationsData, err := pokeapi.GetLocations(cfg.PrevUrl)
    if err != nil {
        return fmt.Errorf("Error returning location data")
    }
    for _, location := range locationsData.Results {
        fmt.Println(location.Name)
    }

    cfg.NextUrl = locationsData.Next
    cfg.PrevUrl = locationsData.Prev

    return nil
}
