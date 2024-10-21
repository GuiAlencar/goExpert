package main

import "fmt"

func main() {
	//memória -> endereço -> valor
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	*b = 30

	fmt.Println(&a)
	fmt.Println(ponteiro)
	fmt.Println(*ponteiro)
	fmt.Println("a -> ", a)
	fmt.Println("b -> ", *b)
}
