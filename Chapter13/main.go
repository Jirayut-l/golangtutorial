package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println( // "bbaa"
		strings.Replace("aaaa", "a", "b", 2),
		// []string{"a","b","c","d","e"}
		strings.Split("a-b-c-d-e", "-"),
		// "test"
		strings.ToLower("TEST"),
		// "TEST"
		strings.ToUpper("test"))
}
