package main

import "fmt"

// MÉTODO 1
// func soma(a *int, b int) int {
// 	*a = 50
// 	return *a + b
// }

// MÉTODO 2
func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	// MÉTODO 1
	// minhaVar1 := 10
	// var ponteiro *int = &minhaVar1
	// minhaVar2 := 20
	// soma(ponteiro, minhaVar2)
	// fmt.Println(minhaVar1)

	minhaVar1 := 10
	minhaVar2 := 20
	soma(&minhaVar1, &minhaVar2)
	fmt.Println(minhaVar1)
}
