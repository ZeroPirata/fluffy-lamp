```none
Given two elements a and b, exactly one of the following must be true: a < b a<b, a = b a=b, or a > b a>b ( Law of Trichotomy )
If a < b a<b and b<c b<c, then a<c a<c ( Law of Transitivity )
```

## Notação big O

Classificação da complexidade do algoritmo.

-   o(1) Operação de custo fixo
-   o(n) Operação de custo proporcional ao tamanho da entrada
-   o(log n) Dividir para conquistar

### O(1)

o custo será exatamente o mesmo, independente do tamanho da entrada.

```go
func number(n int) string {
    return fmt.Sprintf("o numero é %d", n)
}
```

### O(n)

O custo é proporcional ao tamanho da entrada.

```go
func number(n int) string {
    for i := 0; i < n; i++ {
        fmt.Println(i)
    }
    return fmt.Sprintf("o numero é %d", n)
}
```

### O(n^2)

O custo é proporcional ao quadrado do tamanho da entrada.

```go
func number(n int) string {
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Println(i, j)
        }
    }
    return fmt.Sprintf("o numero é %d", n)
}
```

## Comparison Based Sort

Algoritmos de ordenação baseados em comparação são aqueles que determinam a ordem dos elementos através de comparações diretas entre eles. A ordenação é definida pela relação de comparação estabelecida (por exemplo, menor que, maior que).

## Selection Sort

O Selection Sort é um algoritmo de ordenação com complexidade de tempo O(n^2), o que o torna ineficiente para grandes conjuntos de dados. Ele funciona encontrando repetidamente o menor elemento da parte não ordenada da lista e colocando-o no início da parte ordenada.

```go
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}
```

## Bubble Sort

É um algoritmo de ordenação simples, mas ineficiente para grandes conjuntos de dados, com uma complexidade de tempo de O(n^2) no pior e caso médio. Ele funciona comparando pares de elementos adjacentes e trocando-os se estiverem na ordem errada, repetindo o processo até que nenhum elemento precise ser trocado, indicando que a lista está ordenada.

```go
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
```

## Insertion Sort

Dada uma coleção de inteiros, o Insertion Sort organiza a lista processando-a do início ao fim. A cada passo, ele pega o elemento atual e o compara com os elementos anteriores já ordenados. Se o elemento atual estiver fora de ordem, ele é trocado repetidamente com o elemento anterior até que encontre sua posição correta na parte já ordenada da lista.

Em termos de desempenho, o pior cenário para o Insertion Sort ocorre quando a lista de entrada está em ordem decrescente, pois cada elemento precisará ser movido para o início da sublista ordenada. O melhor cenário, por outro lado, é quando a lista já está ordenada, resultando em um desempenho quase linear.

```text
In what cases does insertion sort preferable? (Select all that appropriately):

    For small inputs, insertion sort almost always outperforms other sorts. Also when the number of inversions is small (the list is almost sorted), insertion sort is quite efficient, since there aren't many insertion operations required.
```

```go
func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}
```
