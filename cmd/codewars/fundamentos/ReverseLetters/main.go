package main

import (
	"fmt"
	"unicode"
)

// Given a string str, reverse it and omit all non-alphabetic characters
// For str = "krishan", the output should be "nahsirk".
// For str = "ultr53o?n", the output should be "nortlu".
// A string consists of lowercase latin letters, digits and symbols.

// v1
func ReverseLetters(s string) string {
	sRune := []rune(s)
	result := make([]rune, 0, len(sRune)/2)
	for _, r := range sRune {
		if unicode.IsLetter(r) {
			result = append(result, r)
		}
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

// v2
func ReverseLettersV2(s string) string {
	var result []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			result = append([]rune{r}, result...)
		}
	}
	return string(result)
}

func main() {
	fmt.Println(ReverseLetters("krishan"))    //dotest("","nahsirk")
	fmt.Println(ReverseLetters("ultr53o?n"))  //dotest("","nortlu")
	fmt.Println(ReverseLetters("ab23c"))      //dotest("","cba")
	fmt.Println(ReverseLetters("krish21an"))  //dotest("","nahsirk")
	fmt.Println(ReverseLetters("ZeroPirata")) //dotest(, "")
}
