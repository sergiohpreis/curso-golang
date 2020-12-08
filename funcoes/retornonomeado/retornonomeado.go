package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	/*
		Como as variaveis de retorno já foram nomeadas, não precisamos
		retornar os valores, podemos usar o "retorno limpo"
	*/
	return // retorno limpo
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)
}
