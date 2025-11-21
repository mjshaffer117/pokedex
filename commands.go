// pokedex
// List of valid functions within the pokedex cli tool

package main

import (
    "fmt"
    "os"
)


func commandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}


func commandHelp() error {
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
