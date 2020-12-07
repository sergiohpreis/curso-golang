package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		Ao definir um array, devemos definir seu tamanho na inicialização
		(ou indicar para o compilador fazer isso)
		Ao definir um slice, não definimos o tamanho, tem tamanho váriavel
	*/
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	/*
		O slice é uma pedaço de um array
		Podemos ter um array de 10 posições, com diversos slices referenciados a ele
		Um slice com as posições 1 e 2, outro com as posições 4,5 e 6 e etc

		Você pode imaginar um slice como:
		Tamanho e um ponteiro para um elemento de um array
	*/
	a2 := [5]int{1, 2, 3, 4, 5}

	// slice: pega do índice 1 ao índice 3, não incluindo o índice 3
	s2 := a2[1:3] // [2, 3]
	fmt.Println(a2, s2)

	// slice: pega do ínicio do array ao índice 2, não incluindo o índíndiceice 2
	s3 := a2[:2] // [1,2]
	fmt.Println(a2, s3)

	// podemos fazer um slice de um slice também
	s4 := s2[:1] // [2]
	fmt.Println(s2, s4)
}
