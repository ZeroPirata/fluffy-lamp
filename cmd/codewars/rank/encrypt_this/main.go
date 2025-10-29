package main

import (
	"fmt"
	"strings"
)

// Encrypt this!

// Your message is a string containing space separated words.
// You need to encrypt each word in the message using the following rules:
// The first letter must be converted to its ASCII code.
// The second letter must be switched with the last letter
// Keepin' it simple: There are no special characters in the input.

func main() {
	fmt.Println(EncryptThis("A wise old owl lived in an oak"))
	fmt.Println(EncryptThis("The more he saw the less he spoke"))
	fmt.Println(EncryptThis("The less he spoke the more he heard"))
	fmt.Println(EncryptThis("Why can we not all be like that wise old bird"))
	fmt.Println(EncryptThis("Thank you Piotr for all your help"))
}

// 0123
// wise > 119esi
// 119  > primeira letra		0  [0]
// e    > ultima letra			1  [len(word)-1]
// s    > meio					2  [1:len(word)-1]
// i    > primeira letra		3  [1]

func EncryptThis(text string) string {
	strArray := strings.Split(text, " ")
	for i, word := range strArray {
		if len(word) == 0 {
			continue
		}
		ascii := fmt.Sprintf("%d", word[0])
		switch len(word) {
		case 1:
			strArray[i] = ascii
		case 2:
			strArray[i] = ascii + string(word[1])
		default:
			strArray[i] = ascii + string(word[len(word)-1]) + word[2:len(word)-1] + string(word[1])
		}
	}
	return strings.Join(strArray, " ")
}
