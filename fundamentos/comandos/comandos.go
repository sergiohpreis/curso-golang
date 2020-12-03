// go build: compila o programa
// go run: compila e executa o programa
// go vet: valida o código
package main

import "fmt"

func main() {
	// go vet vai disparar um warning pois esta faltando o arg
	// fmt.Printf("Outro programa em %s!")
	fmt.Printf("Outro programa em %s!\n", "Go")
}
