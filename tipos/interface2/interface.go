package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// Quando implementamos a interface, precisamos passar
	// um ponteiro
	// É mais comum que no nivel de interface, não geremos alteração
	// no objeto
	// Melhor trabalhar com o interfaces pura

	/*
		Ao implementar uma interface diretamente (var nomedaVar tipodaInterface),
		precisamos passar uma referência de ponteiro, caso contrário iremos tomar um erro:
		"
		cannot use ferrari literal (type ferrari) as type esportivo in assignment: ferrari does
		not implement esportivo (ligarTurbo method has pointer receiver)
		"

		Não é muito comum trabalharmos com interfaces que geram esse tipo de efeito colateral,
		preferimos sempre trabalhar com interfaces puras
	*/
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
