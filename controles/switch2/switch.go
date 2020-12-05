package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// Nessa forma de usar o switch, ele procura o primeiro case quer der true
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa Tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
