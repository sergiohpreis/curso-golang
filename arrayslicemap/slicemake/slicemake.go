package main

import "fmt"

func main() {
	// Internamente cria o array interno associado (10)
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s, len(s), cap(s)) // len = 10, cap = 10

	// Podemos dizer a capacidade do array interno associado (20)
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s)) // len = 10, cap = 20

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s)) // len = 20, cap = 20

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s)) // len = 21, cap = 40
}
