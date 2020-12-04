package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números Inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	// byte = alias para uint8
	fmt.Println("")
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("")
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("")
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	// podemos converter um float64 para um float32, assim também:
	var y = float32(49.99)
	fmt.Println("")
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo de y é", reflect.TypeOf(y))
	// por padrão, o literal (com ponto fluante) é do tipo float 64
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// boolean
	// não aceita números inteiros, ex: var bo bool = 1
	bo := true
	fmt.Println("")
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string (somente aspas duplas)
	s1 := "Olá meu nome é Sergio"
	fmt.Println("")
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Sergio`
	fmt.Println("")
	fmt.Println("O tamanho da string é", len(s2))

	// char??
	// não existe o tipo char
	char := 'a'
	fmt.Println("")
	// O valor é um int32 (tabela unicode)
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}
