package main

import "fmt"

// Write a function that returns both the minimum and maximum number of the given list/array.
// [1,2,3,4,5] --> [1,5]
// [2334454,5] --> [5,2334454]
// [1]         --> [1,1]

// All arrays or lists will always have at least one element,
// so you don't need to check the length. Also, your function will
// always get an array or a list, you don't have to check for null, undefined or similar.

func MinMax(arr []int) [2]int {
	min, max := arr[0], arr[0]
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return [2]int{min, max}
}

func main() {
	fmt.Println(MinMax([]int{1, 2, 3, 4, 5}))
	fmt.Println(MinMax([]int{2334454, 5}))
}
