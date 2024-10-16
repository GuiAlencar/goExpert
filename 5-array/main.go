package main

import "fmt"

func main() {
	var meuarray [3]int
	//meuarray2 := [2]int{1, 2}
	meuarray[0] = 1
	meuarray[1] = 2
	meuarray[2] = 10

	fmt.Println(meuarray[0])
	fmt.Println(len(meuarray))
	//pegando ultimo elemento do array
	fmt.Println(meuarray[len(meuarray)-1])

	for indice, valor := range meuarray {
		fmt.Printf("O valor do índice é %d e o valor é %d \n", indice, valor)
	}

}
