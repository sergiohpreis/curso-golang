package main

import "fmt"

func main() {
	i := 1 //ocupa 64 bits na memoria (8 bytes)

	// Ponteiro = Referência de memória, compartilhada entre referencias
	// Ao mudar a referência, todos as referências recebem a alteração

	var p *int = nil // nil = null
	p = &i           // pega o endereço da variavel i e atribui para o ponteiro
	*p++             // desreferenciando (pegando o valor) do ponteiro
	i++

	// Go não tem aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
