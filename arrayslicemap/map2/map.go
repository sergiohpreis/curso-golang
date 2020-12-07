package main

import "fmt"

func main() {
	// Inicializando o mapa com valores iniciais
	funcsESalarios := map[string]float64{
		"José João": 11325.45,
		"Pedro":     15456.45,
		"Carlos":    1200.45,
	}

	// Caso exista a chave, substitui, se não existir, cria
	funcsESalarios["Rafael Filho"] = 1350.0
	funcsESalarios["Pedro"] = 1350.0

	// Podemos excluir um elemento que não existe
	delete(funcsESalarios, "Inexistente")

	// for chave, valor
	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
