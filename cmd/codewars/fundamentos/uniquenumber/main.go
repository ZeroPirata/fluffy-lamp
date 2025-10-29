package main

import "fmt"

// There is an array with some numbers. All numbers are equal except for one. Try to find it!

// findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
// findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55
func FindUniq(arr []float32) float32 {
	mapNumbers := make(map[float32]int)
	for i := range arr {
		mapNumbers[arr[i]]++
	}
	for key, value := range mapNumbers {
		if value == 1 {
			return key
		}
	}
	return arr[0]
}

func main() {
	fmt.Println(FindUniq([]float32{1, 1, 1, 2, 1, 1}))
	fmt.Println(FindUniq([]float32{0, 0, 0.55, 0, 0}))
}
