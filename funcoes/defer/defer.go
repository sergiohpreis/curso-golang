package main

import "fmt"

/*
	É comum usar o "defer" para fechar algum recurso aberto, por exemplo,
	uma conexão com o banco de dados
*/

func obterValorAprovado(numero int) int {
	// Atrasa a execução para o ultimo momento nesse caso antes do "return"
	defer fmt.Println("fim!")

	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	/*
		Log:
		Valor alto...
		fim!
		5000
	*/
	fmt.Println(obterValorAprovado(6000))
	/*
		Log:
		Valor baixo...
		fim!
		3000
	*/
	fmt.Println(obterValorAprovado(3000))
}
