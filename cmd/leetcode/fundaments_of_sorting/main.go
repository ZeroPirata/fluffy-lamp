package main

import "fmt"

// Bubble sort
func SortByLength(arr []string) []string {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if SortingsWords(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func SortNumbers(arr []int) []int {
	n := len(arr)
	trocas := 0
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				trocas++
			}
		}
	}
	fmt.Println(trocas)
	return arr
}

// Essa função é responasvel por fazer a ordenação das palabras, ela faz a comparação de duas strings e seus tamanhos
// caso s1 seja maior que s2 ele retorna true, caso contario ele retorna false que é o final da operação.
func SortingsWords(s1, s2 string) bool {
	return len(s1) > len(s2)
}

// Ordenação utilizando a conotação O(n^2) onde ele utiliza dois for para fazer a ordenação
// ele precisa da váriavel minIndex onde ele armazena o indice de menor indice da comparação (i == i + i)
// com base nisso ele faz a subistituição do indice de menor para o indice de maior
// trocando arr[i] == arr[minIndex]
func SelectionSort(arr []int) []int {
	var minIndex int
	n := len(arr)
	for i := range n {
		minIndex = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = tmp
	}
	return arr
}

// BubbleSort
// Optimized BubbleSort
// A otimização do BubbleSort é que ele para de fazer as interações caso o array ja esteja ordenado
// ele usa uma variavel boleana para fazer essa verificação
func BubbleSort(arr []int) []int {
	n := len(arr)
	hasSwapped := true
	for hasSwapped {
		hasSwapped = false
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				hasSwapped = true
			}
		}
	}
	return arr
}

func heightChecker(heights []int) int {
	h := make([]int, len(heights))
	copy(h, heights)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(h)-1; i++ {
			if h[i] > h[i+1] {
				h[i], h[i+1] = h[i+1], h[i]
				swapped = true
			}
		}
	}
	count := 0
	for i := range heights {
		if heights[i] != h[i] {
			heights[i] = h[i]
			count++
		}
	}
	return count
}

func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	arrInt := []int{-1, 5, 3, 4, 0}
	a := insertionSort(arrInt)
	fmt.Println(a)
}
