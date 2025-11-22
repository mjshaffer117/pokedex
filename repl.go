// pokedex
package main

import(
    "os"
    "fmt"
    "strings"
    "bufio"
)


func startRepl(cfg *Config) {
    reader := os.Stdin
    scanner := bufio.NewScanner(reader)
    for true {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        token := scanner.Text()
        input := cleanInput(token)
        if len(input) == 0 {
            continue
        }
        commandName := input[0]
        command, exists := getCommands()[commandName]
        if exists == false {
            fmt.Println("Unknown command")
            continue
        } else {
            err := command.callback(cfg)
            if err != nil {
                fmt.Println(err)
            }
            continue
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Printf("An error occured: %v\n", err)
    }
}


func cleanInput(text string) []string {
    var words []string
        text = strings.TrimSpace(strings.ToLower(text))
        words = strings.Fields(text)

    return words
}


type cliCommand struct {
    name          string
    description   string
    callback      func(cfg *Config) error
}


func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name:        "help",
            description: "Displays the help menu",
            callback:    commandHelp,
        },
        "map": {
            name:        "map",
            description: "Displays the next 20 map locations",
            callback:    commandMap,
        },
        "mapb": {
            name:        "mapb",
            description: "Displays the previous 20 map locations (back button)",
            callback:    commandMapB,
        },
    }
}
