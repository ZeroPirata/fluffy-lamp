package main

import "fmt"

// number of people that get on the bus (the first item)
// number of people that get off the bus (the second item) at a bus stop.

// number of people who are still on the bus after the last bus stop (after the last array).
// Even though it is the last bus stop, the bus might not be empty and some
// people might still be inside the bus, they are probably sleeping there :D

func Number(busStops [][2]int) (lastStop int) {
	for i := 0; i < len(busStops); i++ {
		join, leave := busStops[i][0], busStops[i][1]
		lastStop += join - leave
	}
	return lastStop
}

func main() {
	fmt.Println(Number([][2]int{{10, 0}, {3, 5}, {5, 8}}))
	fmt.Println(Number([][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}))
	fmt.Println(Number([][2]int{{3, 0}, {9, 1}, {4, 8}, {12, 2}, {6, 1}, {7, 8}}))
	fmt.Println(Number([][2]int{{0, 0}}))
}
