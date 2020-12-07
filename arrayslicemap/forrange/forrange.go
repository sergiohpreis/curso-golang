package main

import "fmt"

func main() {
	/*
		Nessa forma de inicialização, quem ira fazer a contagem do tamanho
		do array é o compilador
	*/
	numeros := [...]int{1, 2, 3, 4, 5}

	// Equivalente ao forEach
	for i, numero := range numeros {
		fmt.Printf("%d) %d \n", i+1, numero)
	}

	// Usando "_" para ignorar o indice
	for _, num := range numeros {
		fmt.Println(num)
	}
}
