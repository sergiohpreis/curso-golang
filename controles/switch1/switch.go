package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	/*
		O switch do Go tem o comportamento padrão de executar
		um case, e ja sair, fazendo com que a palavra "break"
		não seja necessária
	*/
	switch nota {
	case 10:
		/*
			Ao invés do "break", usamos o "fallthrough" para indicar que o
			switch deve passar para o proximo case, e não mais executar seu
			comportamento padrão, que seria sair do switch
		*/
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(11.1))
	fmt.Println(notaParaConceito(-5))
}
