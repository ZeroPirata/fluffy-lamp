package main

import "fmt"

// The maximum sum subarray problem consists in finding the maximum sum of a contiguous subsequence in an array or list of integers

// Input:  [-2, 1, -3, 4, -1, 2, 1, -5, 4]
// Result: 6 => [4, -1, 2, 1]
func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	currentSum := numbers[0]
	maxSum := numbers[0]

	for i := 1; i < len(numbers); i++ {
		n := numbers[i]
		currentSum = max(n, currentSum+n)
		maxSum = max(maxSum, currentSum)
	}

	if maxSum < 0 {
		return 0
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))      // 6
	fmt.Println(MaximumSubarraySum([]int{-2, -1, -3, -4, -1, -2, -1, -5, -4})) // 0
}
