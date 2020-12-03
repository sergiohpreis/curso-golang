package main

import "fmt"

func main() {
	// Imprime na mesma linha
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// Imprime quebrando uma linha no final
	fmt.Println(" Nova")
	fmt.Println(" Linha.")

	x := 3.141516

	// Não é possivel concatenar tipos diferentes, ex:
	fmt.Println("O valor de x é " + x)

	// Transforma em string
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)

	// Pode imprimir tipos diferentes usando a virgula
	fmt.Println("O valor de x é", x, "!!!")

	// %.2f = Duas casas decimais do tipo float
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "string"
	// \n = Quebra de linha
	// %d = Inteiro
	// %f = Float
	// %t = Boolean
	// %s = String
	// %v = Genérico
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
