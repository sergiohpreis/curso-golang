package main

import "fmt"

// não tem operador ternário
// return nota >= 6 ? "Aprovado" : "Reprovado"
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}
func main() {
	fmt.Println(obterResultado(6.2))
}
