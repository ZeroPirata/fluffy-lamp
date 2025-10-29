package main

import "fmt"

// The Vigenère cipher is a classic cipher originally developed by Italian cryptographer
// Giovan Battista Bellaso and published in 1553. It is named after a later French cryptographer Blaise de Vigenère,
// who had developed a stronger autokey cipher (a cipher that incorporates the message of the text into the key).
// The cipher is easy to understand and implement, but survived three centuries of attempts to break it,
// earning it the nickname "le chiffre indéchiffrable" ("the unbreakable cipher")

// Assume the key is repeated for the length of the text, character by character.
// Note that some implementations repeat the key over characters only if
// they are part of the alphabet -- this is not the case here.

// The shift is derived by applying a Caesar shift to a character with the corresponding
// index of the key in the alphabet.

// "my secret code i want to secure"  // message
// "passwordpasswordpasswordpasswor"  // key

// alphabet = "abcdefghijklmnopqrstuvwxyz"
// key      = "passwordpasswordpasswordpa"

// "codewars" --> encode -->  "rovwsoiv"
// "laxxhsj"  --> decode -->  "waffles"

type VigenèreCipher struct {
	Key   string
	Alpha string
}

func main() {
	cipher := VigenèreCipher{
		Key:   "password",
		Alpha: "abcdefghijklmnopqrstuvwxyz",
	}

	fmt.Println(cipher.Encode("codewars"), "rovwsoiv") // rovwsoiv
	fmt.Println(cipher.Decode("rovwsoiv"), "codewars") // codewars
	fmt.Println(cipher.Encode("waffles"), "waffles")   // laxxhsj
	fmt.Println(cipher.Decode("laxxhsj"), "laxxhsj")   // waffles
	fmt.Println(cipher.Encode("CODEWARS"), "CODEWARS") // CODEWARS
	fmt.Println(cipher.Decode("CODEWARS"), "CODEWARS") // CODEWARS
	fmt.Println(cipher.Decode("カタカナ"), "CODEWARS")     // CODEWARS

}

// atacar -> password
// p
// a b c d e f g h i j k l m n o p q r s t u v w x y z
// a b c d e f g h i j k l m n o p q r s t u v w x y z
// func que vai ficar responsabilizada para fazer a copia da chave do tamanho
// da mensagem original, assim respeitando a regra.
func (c VigenèreCipher) copyKey(str string) string {
	strRunes := []rune(str)
	keyRunes := []rune(c.Key)
	copyKey := ""

	for i := 0; i < len(strRunes); i++ {
		copyKey += string(keyRunes[i%len(keyRunes)])
	}

	return copyKey
}

// func que vai ser responsabilizada para indentificar em qual indice o caracter está
// caso não ache o caractere ele vai retornar como -1 que é equivalente
// a não encontrar dentro do alph
func (c VigenèreCipher) findIndex(caractere rune) int {
	for i, alphRune := range []rune(c.Alpha) {
		if caractere == alphRune {
			return i
		}
	}
	return -1
}

func (c VigenèreCipher) Encode(str string) string {
	stringEncoded := ""
	strRunes := []rune(str)
	copyKey := c.copyKey(str)
	keyRunes := []rune(copyKey)
	alphaRunes := []rune(c.Alpha)

	for i := 0; i < len(strRunes); i++ {
		char := strRunes[i]
		strIndex := c.findIndex(char)
		if strIndex == -1 {
			stringEncoded += string(char)
			continue
		}

		key := keyRunes[i]
		copyIndex := c.findIndex(key)

		newLatter := (strIndex + copyIndex) % len(alphaRunes)
		stringEncoded += string(alphaRunes[newLatter])
	}

	return stringEncoded
}

func (c VigenèreCipher) Decode(str string) string {
	stringDecoded := ""
	strRunes := []rune(str)
	copyKey := c.copyKey(str)
	keyRunes := []rune(copyKey)
	alphaRunes := []rune(c.Alpha)

	for i := 0; i < len(strRunes); i++ {
		char := strRunes[i]
		strIndex := c.findIndex(char)
		if strIndex == -1 {
			stringDecoded += string(char)
			continue
		}

		key := keyRunes[i]
		copyIndex := c.findIndex(key)

		newLatter := (strIndex - copyIndex + len(alphaRunes)) % len(alphaRunes)
		stringDecoded += string(alphaRunes[newLatter])
	}

	return stringDecoded
}
