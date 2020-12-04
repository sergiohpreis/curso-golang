package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	/*
		Não funciona dividindo direto, por exemplo
		fmt.Println(x / y)
		É necessário uma conversão
	*/
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	nota := 6.9
	// ignora o valor da casa decimal na conversão
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado... dessa forma converte para a tabela unicode (rune)
	fmt.Println("Teste " + string(123))
	// forma correta de converter int para string
	fmt.Println("Teste " + strconv.Itoa(123))
	fmt.Println("Teste " + fmt.Sprint(123))

	// string para int
	// retorna 2 valores
	// primeiro valor: numero
	// segundo valor: erro
	// _ ignora a variavel (pois em go temos que usar var declarada)
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// string para bool
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
	// valor "1" também é verdadeiro
	c, _ := strconv.ParseBool("1")
	if c {
		fmt.Println("Verdadeiro")
	}
}
