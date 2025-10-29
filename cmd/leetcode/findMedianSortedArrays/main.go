package main

import "fmt"

// Given two sorted arrays nums1 and nums2 of size m and n respectively,
// return the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).

// Example 1:
// 	Input: nums1 = [1,3], nums2 = [2]
// 	Output: 2.00000
// 	Explanation: merged array = [1,2,3] and median is 2.

// Example 2:
// 	Input: nums1 = [1,2], nums2 = [3,4]
// 	Output: 2.50000
// 	Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 1. coisa a se fazer é juntar ambos os array
	//
	newArry := make([]int, 0, len(nums1)+len(nums2))
	newArry = append(newArry, nums1...)
	newArry = append(newArry, nums2...)
	// 2. depois que juntar os array é ordenar os numeros na ordem crescente
	for i := 0; i < len(newArry); i++ {
		for j := i + 1; j < len(newArry); j++ {
			if newArry[i] > newArry[j] {
				newArry[i], newArry[j] = newArry[j], newArry[i]
			}
		}
	}
	// 3. Pegar o numero do meio

	fmt.Println(newArry)
	return 0.0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{1, 3, 4}))
}
