// pokecache
// internal package for pokecache

package pokecache

import (
    "sync"
    "time"
)


type Cache struct {
    entries map[string]cacheEntry
    mtx     sync.Mutex
    quit    chan struct{}
}

type cacheEntry struct {
    createdAt time.Time
    val       []byte
}


func NewCache (interval time.Duration) Cache {
    c := Cache {
        entries: make(map[string]cacheEntry),
        quit:    make(chan struct{}),
    }
    go c.reapLoop(interval)
    return c
}


func (c *Cache) reapLoop (interval time.Duration) {
    ticker := time.NewTicker(interval)
    defer ticker.Stop()

    for {
        select {
        case <- ticker.C:
            c.mtx.Lock()
            for key, entry := range c.entries {
                if entry.createdAt.Add(interval).Before(time.Now()) {
                    delete(c.entries, key)
                }
            }
            c.mtx.Unlock()
        case <- c.quit:
            return
        }
    }
}


func (c *Cache) Get (key string) ([]byte, bool) {
    c.mtx.Lock()
    defer c.mtx.Unlock()

    entry, exists := c.entries[key]
    if exists {
        return entry.val, exists
    } else {
        return nil, exists
    }
}


func (c *Cache) Add (key string, val []byte) {
    var newEntry cacheEntry
    newEntry.createdAt = time.Now()
    newEntry.val = val

    c.mtx.Lock()
    defer c.mtx.Unlock()

    c.entries[key] = newEntry
}


func (c *Cache) Stop() {
    close(c.quit)
}
