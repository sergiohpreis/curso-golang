package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)
	/*
		Ao incluir um elemento na posição 0, ele vai incluir tanto para o
		s1 quanto para o s2, pois ambos referenciam o mesmo array internamente
	*/
	s1[0] = 7
	fmt.Println(s1, s2)
}
