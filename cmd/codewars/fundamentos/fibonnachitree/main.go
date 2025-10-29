package main

import "fmt"

func Tribonacci(signature [3]float64, n int) []float64 {
	if len(signature) == 0 || n == 0 {
		return []float64{}
	}
	if n < 3 {
		return signature[:n]
	}
	nums := make([]float64, n)
	nums[0], nums[1], nums[2] = signature[0], signature[1], signature[2]
	for i := 3; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2] + nums[i-3]
	}
	return nums
}

func main() {
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 10))
	fmt.Println(Tribonacci([3]float64{0, 0, 1}, 10))
	fmt.Println(Tribonacci([3]float64{0, 1, 1}, 10))
	fmt.Println(Tribonacci([3]float64{1, 0, 0}, 10))
	fmt.Println(Tribonacci([3]float64{0, 0, 0}, 10))
	fmt.Println(Tribonacci([3]float64{1, 2, 3}, 10))
	fmt.Println(Tribonacci([3]float64{3, 2, 1}, 10))
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 1))
	fmt.Println(Tribonacci([3]float64{300, 200, 100}, 0))
	fmt.Println(Tribonacci([3]float64{0.5, 0.5, 0.5}, 30))
}
