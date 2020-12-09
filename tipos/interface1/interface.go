package main

import "fmt"

/*
	A interface imprimivel declara que toda estrutura
	desse tipo, precisa ter o método "toString()" associado
	a ela
*/
type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

/*
	Não existe marcação direta entre a struct "produto" ou "pessoa" com a interface
	"imprimivel".
	A interface é implementada implicitamente a partir do momento que atribuimos
	o método "toString()" a essas structs, conforme exemplo abaixo:
*/

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s = R$ %.2f", p.nome, p.preco)
}

/*
	Nesse caso, x pode ser tanto do tipo "pessoa" quando do tipo "produto"
	pois ambos implementam o metodo "toString()"
*/
func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	/*
		Nos exemplos acima, funciona como uma espécie de "polimorfismo", pois a var
		coisa, em um determinado momento, é do tipo "pessoa", e em outro é do tipo
		"produto".

		Isso é muito bom pois é preferível composição a herança
	*/

	p2 := produto{"Calça Jeans", 179.90}
	imprimir(p2)
}
