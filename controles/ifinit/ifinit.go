package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	/*
		Bloco de inicialização
		if inicialização; condição
		O escopo da váriavel i somente para o bloco do if
		ou do else
	*/
	if i := numeroAleatorio(); i > 5 { // Também funciona no switch
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!")
	}

	// var i undefined
	// fmt.Println(i)
}
