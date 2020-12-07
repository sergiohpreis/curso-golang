package main

import "fmt"

func main() {
	// Maps devem ser inicializados
	// var aprovados map[int]string
	aprovados := make(map[int]string)

	// Para atribuir um valor a uma chave:
	// map[chave] = valor
	aprovados[12345678900] = "Maria"
	aprovados[12345678901] = "Pedro"
	aprovados[12345678902] = "Ana"
	fmt.Println(aprovados)

	// for chave, valor
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// Para acessar um valor de uma chave:
	// mapa[chave]
	fmt.Println(aprovados[12345678900])

	// Para deletar um elemento do map
	// delete(mapa, chave)
	delete(aprovados, 12345678901)
	fmt.Println(aprovados[12345678901])
}
