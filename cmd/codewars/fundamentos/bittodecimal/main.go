package main

import "fmt"

func BinToDec(bin string) int {
	total := 0
	for i := 0; i < len(bin); i++ {
		total = total*2 + int(bin[i]-'0')
	}
	return total
}

func main() {
	fmt.Println(BinToDec("1"))
	fmt.Println(BinToDec("10"))
	fmt.Println(BinToDec("0"))
	fmt.Println(BinToDec("11"))
	fmt.Println(BinToDec("101010"))
	fmt.Println(BinToDec("1001001"))
}
