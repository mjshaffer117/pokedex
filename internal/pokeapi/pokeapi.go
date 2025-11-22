// pokedex
// api functionality

package pokeapi

import(
    "encoding/json"
    "fmt"
    "net/http"
)


type LocationResponse struct {
    Next string `json:"next"`
    Prev string `json:"previous"`
    Results []LocationArea `json:"results"`
}

type LocationArea struct {
    Name string `json:"name"`
}


func GetLocations(url string) (LocationResponse, error) {
    res, err := http.Get(url)
    if err != nil {
       fmt.Println("Error fetching api data")
       return LocationResponse{}, err
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return LocationResponse{}, fmt.Errorf("Api request failed with status code %d", res.StatusCode)
    }

    var locationData LocationResponse
    err = json.NewDecoder(res.Body).Decode(&locationData)
    if err != nil {
        return LocationResponse{}, fmt.Errorf("Error decoding JSON response")
    }

    return locationData, nil
}
