package main

import (
	"fmt"
	"time"

	"github.com/sergiohpreis/html"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	/*
		Na estrutura de select, o primeiro canal que tiver resultado
		vai disparar o case correspondente
		Nesse exemplo, a primeira url que responder o título através
		da chamada http, ira ser retornada por esta func
	*/
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	// Caso nenhum channel responda em até 800ms
	case <-time.After(800 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://youtube.com",
		"https://amazon.com",
		"https://google.com",
	)
	fmt.Println(campeao)
}
