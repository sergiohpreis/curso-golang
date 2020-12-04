package main

import "fmt"

func main() {
	var a int
	// precisa declarar o tipo, porque se nao o compilador nao consegue inferir, Go é fortemente tipado
	// os tipos são cravados na variavel daquele escopo, não é possível alterar o tipo
	// por isso, ao não atribuir valor, o tipo precisa ser declarado
	var b float64
	var c bool
	var d string
	// ponteiro de inteiro
	var e *int

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}
