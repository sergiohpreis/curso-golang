package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	fmt.Println("Inicio da Execução")
	ch <- 1
	fmt.Println("Adicionado 1 ao Buffer")
	ch <- 2
	fmt.Println("Adicionado 2 ao Buffer")
	ch <- 3
	fmt.Println("Adicionado 3 ao Buffer")
	ch <- 4
	fmt.Println("Adicionado 4 ao Buffer")
	ch <- 5
	fmt.Println("Adicionado 5 ao Buffer")
	ch <- 6
	fmt.Println("Adicionado 6 ao Buffer")
}

func main() {
	// Buffer de 3
	ch := make(chan int, 3)

	go rotina(ch)

	/*
		Nesse caso, o valor 4 da função "rotina" não será enviado ao canal por conta
		do "time.Sleep".
		Como a função vai esperar 1 segundo, no momento da inserção do valor 4, o
		buffer ainda vai estar cheio (contendo os ints 1, 2 e 3)

		Dependendo de quem finalizar primeiro (main ou rotina), pode ser que o dado de valor
		4 seja inserido, se tirarmos o "time.sleep" o dado 4 será inserido com certeza
	*/
	time.Sleep(time.Second)
	fmt.Printf("Dado do Canal > %d", <-ch)
}
