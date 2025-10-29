package main

import (
	"fmt"
	"strings"
)

// The goal of this exercise is to convert a string to a new string where
// each character in the new string is "(" if that character appears only once in the original string,
// or ")" if that character appears more than once in the original string.
// Ignore capitalization when determining if a character is a duplicate.

func main() {
	fmt.Println(DuplicateEncode("din"))
	fmt.Println(DuplicateEncode("recede"))
}

// Se a palavra não existir dentro é colocado "("
// se a palavra existir dentro é colocado ")" em todas
// "din"      =>  "((("
// "recede"   =>  "()()()"
// "Success"  =>  ")())())"
func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	charCounts := make(map[rune]int)
	for _, char := range word {
		charCounts[char]++
	}

	var resultBuilder strings.Builder
	for _, char := range word {
		if charCounts[char] > 1 {
			resultBuilder.WriteRune(')')
		} else {
			resultBuilder.WriteRune('(')
		}
	}
	return resultBuilder.String()
}
