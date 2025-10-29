package main

import "fmt"

// A beer can pyramid will square the number of cans in each level - 1 can in the top level, 4 in the second,
// 9 in the next, 16, 25...

// Complete the beeramid function to return the number of complete levels of a beer can pyramid you can make, given the parameters of:
// your referral bonus, and
// the price of a beer can

// 9 bonus
// 2 preço de cada cerveja

// 1 -> 9 - 2 -> 7 ( 1 * 2 = 1)
// 2 -> 7 - 2 -> 5 ( 2 * 2 = 4)
// 3 -> 5 - 2 -> 3 ( 2 * 2 = 4)
// 4 -> 3 - 2 -> 1 ( 2 * 2 = 4)
func Beeramid(bonus int, price float64) int {
	if bonus < 0 {
		return 0
	}
	bonusFloat := float64(bonus)

	piramidyHeight := 1
	for bonusFloat >= price {
		completeFloor := false
		squareNumber := piramidyHeight * piramidyHeight
		for i := range squareNumber {
			bonusFloat -= price
			if i == squareNumber {
				completeFloor = true
			}
		}
		if completeFloor {
			piramidyHeight++
		}
	}
	return piramidyHeight
}

func main() {
	fmt.Println(Beeramid(9, 2.0))  // 1
	fmt.Println(Beeramid(21, 1.5)) // 3
}