package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	// Nesse caso, é um tipo "interface" genérico
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	/*
		Aqui por exemplo, podemos mudar o valor que uma
		variavel desse tipo recebe sem nenhum problema
	*/
	type dinamico interface{}

	var coisa2 dinamico = "Opa" // string
	fmt.Println(coisa2)

	coisa2 = true // bool
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"} //struct
	fmt.Println(coisa2)
}
