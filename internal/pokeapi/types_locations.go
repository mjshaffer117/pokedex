// pokedex
// internal package

// Struct for location data

package pokeapi

type RespShallowLocations struct {
    Count    int     `json:"count"`
    Next     *string `json:"next"`
    Previous *string `json:"previous"`
    Results []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}
