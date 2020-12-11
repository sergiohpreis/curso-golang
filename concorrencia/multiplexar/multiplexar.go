package main

import (
	"fmt"

	"github.com/sergiohpreis/html"
)

// origem: read-only
// destino: vai receber informações
func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.Titulo("https://www.amazon.com", "https://www.youtube.com"),
	)
	// O primeiro dos 4 canais
	fmt.Println("Primeiro:", <-c)
	fmt.Println("Segundo:", <-c)
	fmt.Println("Terceiro:", <-c)
	fmt.Println("Quarto:", <-c)
}
