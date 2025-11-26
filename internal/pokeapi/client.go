// pokedex
// internal package

// host client

package pokeapi

import (
    "net/http"
    "time"
    "github.com/mjshaffer117/pokedex/internal/pokecache"
)

// Client
type Client struct {
    cache      pokecache.Cache
    httpClient http.Client
}


// New Client
func NewClient(timeout, cacheInterval time.Duration) Client {
    return Client{
        cache: pokecache.NewCache(cacheInterval),
        httpClient: http.Client{
            Timeout: timeout,
        },
    }
}

