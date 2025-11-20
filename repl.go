// pokedex
package main

import(
    "os"
    "fmt"
    "strings"
    "bufio"
)


func startRepl() {
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
        fmt.Printf("Your command was: %v\n", input[0])
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
