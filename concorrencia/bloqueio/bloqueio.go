package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante

	/*
		As próximas linhas só serão executadas após o dado enviado para o
		channel ter sido lido.

		No caso de um channel com buffer, as proximas linhas só serão executadas
		após o buffer ser esvaziado (depois de cheio)

		No exemplo dessa aula, o canal criado não terá buffer
	*/

	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go rotina(c)

	/*
		Abaixo, teremos uma operação bloqueante, pois ela só vai ser executada
		após o dado ser recebido
	*/
	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Foi lido")

	/*
		Aqui teremos um deadlock, pois não tera mais ninguém que possa enviar dados
		para o canal
	*/
	fmt.Println(<-c)   // deadlock
	fmt.Println("Fim") // não será impresso por conta do deadlock
}
