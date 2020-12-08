package main

import "fmt"

// Bem comum achar funcoes que retornam um valor e um erro
func fatorial(n int) (int, error) {
	switch {
	/*
		Uma solução melhor seria usar tipo "uint" para não ter que validar
		o valor negativo
	*/
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
