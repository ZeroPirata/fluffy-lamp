package main

import "fmt"

func main() {
	fmt.Println(Parse("iiisdosodddddiso"))
}

// 1{i}, 2{i}, 3{i}
// 3{s} => 9
// 9{d} => 8
// 8{o} => [ 8 ]
// 8{s} => 64
// 64{o} => [ 8 64 ]

// i: Increment the value					=> x + 1
// d: Decrement the value					=> x - 1
// s: Square the value						=> x**2
// o: Output the value to a result array	=> [x]

func Parse(data string) []int {
	startValue := 0
	result := []int{}
	runes := []rune(data)
	for i := 0; i <= len(runes)-1; i++ {
		switch string(runes[i]) {
		case "i":
			startValue++
		case "d":
			startValue--
		case "s":
			startValue = startValue * startValue
		case "o":
			result = append(result, startValue)
		}
	}
	return result
}
