// pokedex
// Logic for map commands functionality

package main

import (
    "errors"
    "fmt"
)


func commandMap(cfg *config) error {
    locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
    if err != nil {
        return err
    }

    cfg.nextLocationsURL = locationsResp.Next
    cfg.prevLocationsURL = locationsResp.Previous

    for _, location := range locationsResp.Results {
        fmt.Println(location.Name)
    }

    return nil
}


func commandMapB(cfg *config) error {
    if cfg.prevLocationsURL == nil {
        return errors.New("First page reached, use 'map' to jump to the next page")
    }

    locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
    if err != nil {
        return err
    }

    cfg.nextLocationsURL = locationsResp.Next
    cfg.prevLocationsURL = locationsResp.Previous

    for _, location := range locationsResp.Results {
        fmt.Println(location.Name)
    }

    return nil
}

