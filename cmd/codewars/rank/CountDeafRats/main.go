package main

import (
	"fmt"
	"strings"
)

// Story
// The Pied Piper has been enlisted to play his magical tune and coax all the rats out of town.

// But some of the rats are deaf and are going the wrong way!

// Kata Task
// How many deaf rats are there?

// Legend
// P  = The Pied Piper
// O~ = Rat going left
// ~O = Rat going right
//
// Example
// ex1 ~O~O~O~O P has 0 deaf rats

// ex2 P O~ O~ ~O O~ has 1 deaf rat

// ex3 ~O~O~O~OP~O~OO~ has 2 deaf rats

func main() {
	fmt.Println(CountDeafRats("~O~O~O~O P"))      // 0
	fmt.Println(CountDeafRats("P O~ O~ ~O O~"))   // 1
	fmt.Println(CountDeafRats("~O~O~O~OP~O~OO~")) // 2
}

func runSide(s, defOrientation string) int {
	sum := 0
	for i := 0; i < len(s)-1; i += 2 {
		if s[i:i+2] == defOrientation {
			sum += 1
		}
	}
	return sum
}

func CountDeafRats(town string) int {
	town = strings.ReplaceAll(town, " ", "")
	var left, right string
	for i, p := range town {
		if string(p) == "P" {
			left = town[:i]
			right = town[i+1:]
		}
	}
	r := runSide(left, "O~")
	l := runSide(right, "~O")
	return r + l
}
