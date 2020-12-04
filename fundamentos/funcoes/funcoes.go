package main

import "fmt"

// precisamos declarar os tipos dos params
// se declaramos o retorno, so podemos retornar daquele tipo
func somar(a int, b int) int {
	return a + b
}

// se nao declaramos o tipo do retorno, nao precisamos retornar
func imprimir(valor int) {
	fmt.Println(valor)
}
