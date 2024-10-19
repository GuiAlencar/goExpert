package main

import (
	"fmt"
)

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	//função anônima
	total := func() int {
		return soma(1, 3, 4, 5, 8, 6, 10) * 2
	}()
	fmt.Println(total)

	fmt.Println(soma(1, 3, 4, 5, 8, 6, 10))
}
