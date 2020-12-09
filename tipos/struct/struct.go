package main

import "fmt"

/*
	Struct é um agrupador de dados, que funciona parecido
	com uma classe.
	Sintaxe :

	type "nome da struct" struct {
		atributo tipo
	}
*/
type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// funcao associada a estrutura
// receiver - quer aplicar a  funcao a que receptor
// quem sera o dono da função? (produto.????)

// metodo = funcao com receiver
// dentro da funcao produto referenciado por p ao invés de this

/*
	Em Go usamos receivers, que são funções associadas a uma struct,
	é basicamente o mesmo conceito de métodos, porém não declaramos
	eles dentro da struct, sendo assim, fica bem desacoplada a criação
	dessas funções.
	Sintaxe:

	func (alias_receiver receiver) nomeDaFunção tipoRetorno {

	}
*/
func (p produto) precoComDesconto() float64 {
	// Ao invés de usar this.preco, usamos o alias (nesse caso "p")
	return p.preco * (1 - p.desconto)
}

func main() {
	// Definindo uma struct
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	// Forma reduzida de definir uma struct
	produto2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2, produto2.precoComDesconto())
}
