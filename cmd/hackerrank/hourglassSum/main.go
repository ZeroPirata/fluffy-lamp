package main

import "fmt"

func hourglassSum(arr [][]int32) int32 {
	n := len(arr)
	// Define um ponto de partida mínimo (válido para −9 ≤ x ≤ 9)
	var maxSum int32 = -63
	// A variável `maxSum` é inicializada com -63 porque é o menor valor possível para a soma de uma ampulheta.
	// Cada elemento da matriz varia de -9 a 9. Uma ampulheta tem 7 elementos.
	// O menor valor possível para a soma de 7 elementos é 7 * -9 = -63.
	// Isso garante que qualquer soma de ampulheta válida será maior ou igual a `maxSum`,
	// permitindo que a primeira soma calculada seja corretamente atribuída a `maxSum`.
	for i := 0; i <= n-3; i++ {
		// A lógica dos loops `for i := 0; i <= n-3; i++` e `for j := 0; j <= n-3; j++`
		// é para iterar sobre todas as possíveis posições iniciais (canto superior esquerdo)
		// de uma ampulheta 3x3 dentro da matriz `arr`.
		//
		// `n` é o número de linhas (e colunas, assumindo uma matriz quadrada) da matriz.
		//
		// O `-3` em `n-3` é porque uma ampulheta tem 3 linhas e 3 colunas.
		// Se a matriz tem `n` linhas, a última linha onde uma ampulheta pode começar
		// é `n-3`. Por exemplo, se `n=6`, a ampulheta pode começar na linha 0, 1, 2, 3.
		// Se começasse na linha 4, não haveria linhas 4+1 e 4+2 para completar a ampulheta.
		// O mesmo raciocínio se aplica às colunas com `j <= n-3`.
		for j := 0; j <= n-3; j++ {
			// Calcula a soma dos elementos da ampulheta atual
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] +
				arr[i+1][j+1] +
				arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			// A lógica `if i == 0 && j == 0 || sum > maxSum` garante que `maxSum` seja inicializado corretamente
			// com a soma da primeira ampulheta (quando `i` e `j` são 0) ou atualizado se uma soma maior for encontrada.
			// Isso é crucial porque `maxSum` é inicializado com um valor mínimo (-63), e a primeira soma calculada
			// deve ser considerada como o maior valor inicial para comparação.
			if i == 0 && j == 0 || sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func main() {
	arr := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	r := hourglassSum(arr)
	fmt.Println(r)
}
