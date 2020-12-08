package main

import "fmt"

/*
	É uma convenção em GO que a função "init" sempre vai ser executada
	Ela é bem util para fazer processos de inicializações
*/
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
