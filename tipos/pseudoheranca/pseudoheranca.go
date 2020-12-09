package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	/*
		Ao criar um "campo anonimo", a struct recebe todos os atributos
		desse tipo, nesse caso, ira receber os campos "nome" e "velocidadeAtual"
		da struct "carro"
		Isso não funciona como uma herança, mas sim como uma composição
	*/
	carro       // campos anonimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrar %s esta com turbo ligado? %v \n", f.nome, f.turboLigado)
	// Esse log mostra a formatação, e mostra a composição que é feita
	fmt.Println(f) // {{F40 0} true}
}
