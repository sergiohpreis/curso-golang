package main

import "fmt"

func main() {
	/*
		Em Go, o canal é um tipo, ou seja, é algo intrinseco da
		propria linguagem (não precisando de uma lib por exemplo).

		Um canal pode ser usado como ponto de comunicação entre
		"goroutines", uma vez que ele recebe e transmite dados.
		Ele age como um ponto de sincronismo, através do canal,
		posso esperar os dados de funções independentes chegarem

		Um canal pode gerar "deadlock", por exemplo, se as "goroutines"
		terminarem e o canal continuar esperando o recebimento de dados
	*/

	// Criamos um canal mais ou menos igual array (com make)
	ch := make(chan int, 1) // canal de inteiros, de buffer 1

	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
