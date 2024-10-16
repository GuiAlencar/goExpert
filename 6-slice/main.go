package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10, 20}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
