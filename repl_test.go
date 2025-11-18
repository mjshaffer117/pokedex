// pokedex
package main

import (
    "testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input string
        expected []string
    }{
        {
            input: "  hello world  ",
            expected: []string{"hello", "world"},
        },
        {
            input: "Hello There Person",
            expected: []string{"hello", "there", "person"},
        },
        {
            input: "Charmander Squirtle",
            expected: []string{"charmander", "squirtle"},
        },
    }
    // add more test cases here...

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("Lengths of actual slice --> '%d', and expected slice --> '%d' do not match.", len(actual), len(c.expected))
            return
        }
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("Actual word '%s' does not match expected word '%s'", word, expectedWord)
                return
            }
        }
    }
}
