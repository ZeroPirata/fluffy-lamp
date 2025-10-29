package main

import "fmt"

func main() {
	staircase(6)
}

func staircase(n int32) {
	// Para escrever o " " é necessario ter o numero de corridas - i
	// i é o numero de corridas que vai ser dado durante todo o percuso com base no n
	for i := 0; i <= int(n); i++ {
		if i == 0 {
			continue
		}
		// j é a quantidade de interações que vai acontecer com base em (n - i)
		for range int(n) - i {
			fmt.Print(" ")
		}
		// k é a quantidade de interações que vai acontecer com o restante das corridas de i
		for range i {
			fmt.Print("#")
		}
		// pular para a proxima linha
		fmt.Println()
	}
}
