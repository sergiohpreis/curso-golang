package main

import "fmt"

func main() {
	/*
		Um array em Go é:
		Homogêneos: uma vez definido o tipo não pode ser alterado
		Estático: uma vez definido o tamanho, ele não pode variar

		Ex:
		var notas [3] float64

		Esse array só vai ter 3 elementos do tipo float64
	*/
	var notas [3]float64
	// Nesse caso, é inicializado com o valor padrão do float64 (0)
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	// Se tentarmos colocar um quarto elemento, irá dar erro, ex:
	// notas[3] = 10
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)
}
