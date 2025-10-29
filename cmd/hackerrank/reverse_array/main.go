package main

import "fmt"

func reverseArray(a []int32) []int32 {
	n := len(a)
	// A lógica `for i := 0; i < n/2; i++` itera apenas sobre a primeira metade do array.
	// Isso é feito porque, ao trocar elementos simetricamente do início e do fim do array,
	// cada par de elementos é trocado apenas uma vez.
	// Por exemplo, se o array tem 5 elementos (n=5), `n/2` é 2.
	// O loop executará para i = 0 e i = 1.
	// Quando i = 0, `a[0]` é trocado com `a[4]`.
	// Quando i = 1, `a[1]` é trocado com `a[3]`.
	// O elemento do meio (a[2]) não precisa ser trocado consigo mesmo.
	//
	// A expressão `a[i], a[n-i-1] = a[n-i-1], a[i]` realiza a troca dos elementos.
	// `a[i]` refere-se ao elemento na posição `i` a partir do início do array.
	// `a[n-i-1]` refere-se ao elemento na posição `i` a partir do final do array.
	// Por exemplo, para `i=0`, `n-i-1` é `n-1`, que é o último elemento.
	// Para `i=1`, `n-i-1` é `n-2`, que é o penúltimo elemento, e assim por diante.
	// Esta é uma maneira eficiente de inverter um array no local, sem a necessidade de um array auxiliar.
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
	return a
}

func main() {
	arr := []int32{1, 2, 3}
	reverseArray(arr)
	fmt.Println(arr)
}
