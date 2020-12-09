package main

import "fmt"

type nota float64

/*
	Se tentarmos usar o receiver como tipo "float64" (Ex: func (n float64)),
	iremos ter um erro falando que não é possível definir novos metodos para tipos
	não locais ("cannot define new methods to non local type float64")
*/
func (n nota) entre(inicio, fim float64) bool {
	// convertemos "n" para "float64" pois "n" é do tipo "nota"
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(5.0, 7.99):
		return "C"
	case n.entre(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
