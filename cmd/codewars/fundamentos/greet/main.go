package main

import "fmt"

func greet() string {
	helloWorld := []byte("hello world!")
	novoMundo := []rune("")
	ALL := 255

	for _, value := range helloWorld {
		for i := range ALL {
			if value == byte(i) {
				novoMundo = append(novoMundo, rune(i))
			}
		}
	}

	return string(novoMundo)
}

func main() {
	fmt.Println(greet())
}
