package main

import "fmt"

// Greed is a dice game played with five six-sided dice.
// Your mission, should you choose to accept it, is to score a throw according to these rules.
// You will always be given an array with five six-sided dice values.

// Three 1's => 1000 points
// Three 6's =>  600 points
// Three 5's =>  500 points
// Three 4's =>  400 points
// Three 3's =>  300 points
// Three 2's =>  200 points
// One   1   =>  100 points
// One   5   =>   50 point

// Each of 5 dice can only be counted once in each roll.
// For example, a given "5" can only count as part of a triplet
// (contributing to the 500 points) or as a single 50 points, but not both in the same roll.

// Example scoring

//  Throw       Score
//  ---------   ------------------
//  5 1 3 4 1   250:  50 (for the 5) + 2 * 100 (for the 1s)
//  1 1 1 3 1   1100: 1000 (for three 1s) + 100 (for the other 1)
//  2 4 4 5 4   450:  400 (for three 4s) + 50 (for the 5)

func main() {
	fmt.Println("Score for [2, 3, 4, 6, 2]:", Score([5]int{2, 3, 4, 6, 2})) // Expected: 0
	fmt.Println("Score for [4, 4, 4, 3, 3]:", Score([5]int{4, 4, 4, 3, 3})) // Expected: 400
	fmt.Println("Score for [2, 4, 4, 5, 4]:", Score([5]int{2, 4, 4, 5, 4})) // Expected: 450
	fmt.Println("Score for [5, 1, 3, 4, 1]:", Score([5]int{5, 1, 3, 4, 1})) // Expected: 250
	fmt.Println("Score for [1, 1, 1, 3, 1]:", Score([5]int{1, 1, 1, 3, 1})) // Expected: 1100
	fmt.Println("Score for [1, 1, 1, 1, 1]:", Score([5]int{1, 1, 1, 1, 1})) // Expected: 1200
	fmt.Println("Score for [5, 5, 5, 5, 5]:", Score([5]int{5, 5, 5, 5, 5})) // Expected: 600
	fmt.Println("Score for [2, 2, 3, 3, 6]:", Score([5]int{2, 2, 3, 3, 6})) // Expected: 0
	fmt.Println("Score for [3, 3, 3, 3, 3]:", Score([5]int{3, 3, 3, 3, 3})) // Expected: 300
}

func Score(dice [5]int) int {
	counts := make(map[int]int)
	for _, die := range dice {
		counts[die]++
	}

	fmt.Println(counts)

	score := 0
	tripletScores := map[int]int{1: 1000, 2: 200, 3: 300, 4: 400, 5: 500, 6: 600}
	for number, count := range counts {
		if count >= 3 {
			score += tripletScores[number]
			count -= 3
		}
		switch number {
		case 1:
			score += count * 100
		case 5:
			score += count * 50
		}
	}
	return score
}
