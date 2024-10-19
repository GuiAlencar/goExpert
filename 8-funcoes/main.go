package main

import (
	"errors"
	"fmt"
)

func soma(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}

func main() {
	valor, err := soma(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}
