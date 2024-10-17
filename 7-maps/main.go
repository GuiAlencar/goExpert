package main

import "fmt"

func main() {
	//              chave  valor
	salarios := map[string]int{
		"Guilherme": 10000,
		"Henrique":  12000,
		"Alencar":   15000,
	}

	fmt.Println(salarios["Guilherme"])
	delete(salarios, "Guilherme")
	salarios["Gui"] = 5000
	fmt.Println(salarios["Gui"])

	sal := make(map[string]int)
	sal1 := map[string]int{}
	sal1["Wesley"] = 1000
	sal["teste"] = 2000

	fmt.Println(sal1["Wesley"])
	fmt.Println(sal["teste"])

	for chave, valor := range salarios {
		fmt.Printf("O salário do %s é de %d \n", chave, valor)
	}

	for _, valor := range salarios {
		fmt.Printf("O salário é de %d \n", valor)
	}
}
