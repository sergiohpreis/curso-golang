package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
	A função que sera usada como handle no método "http.HandleFunc"
	precisa receber 2 paramêtros (mesmo que não usados)

	- http.ResponseWriter: Responsável por escrever na resposta
	- *http.Request: Um ponteiro para a requisição em si
*/
func horaCerta(w http.ResponseWriter, r *http.Request) {
	/*
		A formatação em Go ocorre diferente de outras linguagens.
		A string "02/01/2006 03:04:05" representa:
		- 02 equivale ao dia
		- 01 equivale ao mês
		- 2006 equivale ao ano
		- 03 equivale a hora
		- 04 equivale ao minuto
		- 05 equivale aos segundos
	*/
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Hora certa %s</h1>", s)
}

func main() {
	// Função que lida com a requisição
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
