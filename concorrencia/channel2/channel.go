package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base
	time.Sleep(time.Second * 3)
	c <- 4 * base
}

func main() {
	/*
		Enquanto uma "goroutine" causa assincronismo, um "channel"
		casa um sincronismo.
	*/
	c := make(chan int)
	go doisTresQuatroVezes(2, c) // assíncrono
	fmt.Println("A")             // impresso imediatamente

	a, b := <-c, <-c  // recebendo dados do canal (sincrono)
	fmt.Println("B")  // impresso após os dados chegarem no canal
	fmt.Println(a, b) // impresso do PrintLn acima

	fmt.Println(<-c) // impresso após o terceiro dado chegar no canal

	/*
		Se tentarmos obter mais um dado do canal, iremos causa um deadlock:
		fatal error: all goroutines are asleep - deadlock!
	*/
	// fmt.Println(<-c)
}
