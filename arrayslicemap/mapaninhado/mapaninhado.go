package main

import "fmt"

func main() {
	// Um map cujo a chave é uma string e o valor é um outro map [string]float64
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gustavo":  1200.45,
			"Gabriela": 13541.00,
		},
		"J": {
			"José":   7634.38,
			"Josias": 121341.12,
		},
		"P": {
			"Pedro": 578.38,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		for nome, valor := range funcs {
			fmt.Println(nome, valor)
		}
	}
}
