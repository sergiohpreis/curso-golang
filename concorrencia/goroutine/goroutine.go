package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%d) %s: %s\n", i+1, pessoa, texto)
	}
}

func main() {
	/*
		Nesse primeiro caso, as funções são executadas de forma
		sequencia, ou seja, irá imprimir 3x a frase da "Maria"
		depois a frase do João
	*/
	// fale("Maria", "Pq vc não fala comigo?", 3)
	// fale("João", "Só posso falar depois de vc!", 1)

	/*
		Ao usar a palavra reservada "go" na frente de uma função, estamos
		criando uma "goroutine", que são funções executadas de forma
		independente, como se fosse uma thread.
		Se criassemos somente 2 "goroutines", assim:

		// go fale("Maria", "Ei...", 500)
		// go fale("João", "Opa...", 500)

		O programa finalizaria sem executar nada, pois a função "fale" tem um
		"sleep" de 1 segundo.
		As "goroutines" só são executadas enquanto a thead principal (main) estiver
		aberta, nesse caso, poderiamos fazer elas serem chamadas assim:

		// go fale("Maria", "Ei...", 500)
		// go fale("João", "Opa...", 500)
		// time.Sleep(time.Second * 5)
		// fmt.Println("Fim!")

		Pois a linha de execução principal vai aguardar 5 segundos antes de ser
		finalizada (por conta do time.Sleep)
	*/

	/*
		Nesse exemplo, podemos ver que a ordem em que as falas da "Maria" são
		executadas não é sempre a mesma, isso porque a execução dela esta rodando
		em paralelo com a execução da fala de "João"
	*/
	go fale("Maria", "Entendi!!", 10)
	fale("João", "Parabéns!", 5)
}
