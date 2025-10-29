package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	x := OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123")
	fmt.Println(x)
}

// Ele vai somar os numeros e ordernar pelo seu valor
// exemplo:
// 56 65 74 100 99 68 86 180 90
// 11 11 11   1 18 14 14   9  9
// A ordem, caso seja o mesmo numero vai continuar a mesma
// do menor para o maior
//
// resultado:
// 100 180 90 56 65 74 68 86 99

type Item struct {
	Original string
	Weight   int
}

func OrderWeight(strng string) string {
	if strings.TrimSpace(strng) == "" {
		return ""
	}
	stringNumbers := strings.Fields(strng)

	items := make([]Item, len(stringNumbers))
	for i, sNum := range stringNumbers {
		sum := 0
		for _, char := range sNum {
			digit, _ := strconv.Atoi(string(char))
			sum += digit
		}
		items[i] = Item{Original: sNum, Weight: sum}
	}

	n := len(items)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if items[i].Weight > items[j].Weight || (items[i].Weight == items[j].Weight && items[i].Original > items[j].Original) {
				items[i], items[j] = items[j], items[i]
			}
		}
	}

	resultStrings := make([]string, n)
	for i, item := range items {
		resultStrings[i] = item.Original
	}

	return strings.Join(resultStrings, " ")
}
