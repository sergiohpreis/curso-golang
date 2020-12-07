package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	/*
		Não podemos usar append a partir de um array
		somente a partir de slices, por exemplo:

		array1 = append(array1, 4, 5, 6)
	*/

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	/*
		A função copy não expande o tamanho do slice
		nesse caso, irá copiar os 2 primeiros elementos pois o slice2 tem
		tamanho de 2
	*/
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2) // [4,5]

	/*
		Não podemos copiar a partir de um array, por exemplo:

		copy(slice2, array1)
	*/
}
