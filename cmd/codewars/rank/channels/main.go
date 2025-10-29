package main

import (
	"fmt"
	"sync"
)

func main() {
	// Aqui eu tenho que declarar o inicio das go routines
	// utilizando a biblioteca sync.WaitGroup
	var wg sync.WaitGroup
	// Aqui eu crio um canal de comunicação utilizando chan, onde é obrigatorio escrever
	// com um inteiro
	odd := make(chan int)
	even := make(chan int)

	// Criação da goRoutine para a execução do primeiro processo de fibonacchi dentro da aplicação
	// a função close(c) é para finalizar a comunicação entre eles
	wg.Go(func() {
		fibonacci(50, odd, even)
		close(odd)
		close(even)
	})

	// Aqui acontece a leitura do canal.
	wg.Go(func() {
		for n := range even {
			fmt.Println("Número par:", n)
		}
	})

	wg.Go(func() {
		for n := range odd {
			fmt.Println("Número impar:", n)
		}
	})

	// Espera todas as go routines acontecer para finalmente finalizar o processo
	wg.Wait()
}

// Função do fibonnachi normal, onde ela inicia com 0 e 1 e vai somando até o numero maximo de interações.
func fibonacci(sequence int, odd, even chan int) {
	nums := make([]int64, sequence)
	nums[0], nums[1] = 0, 1
	for i := 2; i < sequence; i++ {
		nums[i] = nums[i-1] + nums[i-2]
		if nums[i]%2 == 0 {
			even <- int(nums[i])
		} else {
			odd <- int(nums[i])
		}
	}
}
