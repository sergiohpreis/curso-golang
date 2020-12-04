package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix
	// prefix = --y
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // y -= 1 ou y = y - 1
	fmt.Println(y)

	// Não pode decrementar dentro de uma comparação, ex:
	// fmt.Println(x == y--)
}
