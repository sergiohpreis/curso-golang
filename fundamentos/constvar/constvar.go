package main

// m antes do math é um alias para o import
import (
	"fmt"
	m "math"
)

func main() {
	//const nomedaconst tipodaconst = valordaconst
	const PI float64 = 3.1415
	// tipo (float64) inferido pelo compilador
	var raio = 3.2

	//forma reduzida de criar uma var
	// : var ja foi definida
	// := var nao existe, declara valor e inicializa (preferencialmente usar essa)
	// math.pow() = potencia
	// P.S: nao podemos declarar var e nao usar em GO
	area := PI * m.Pow(raio, 2)
	// O valor da var pode ser reatribuido, ex:
	// area = 1
	fmt.Println("A área da circunferência é", area)

	// declarar mais de uma const
	const (
		a = 1
		b = 2
	)

	// declarar mais de uma var
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// var e e f do tipo boolean
	// e = true
	// f = false
	var e, f bool = true, false
	fmt.Println(e, f)

	// g = 2,
	// h = false
	// i = "texto"
	g, h, i := 2, false, "texto"
	fmt.Println(g, h, i)

	// o tipo das var nao variam durante o programa
}
