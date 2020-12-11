package main

import "math"

/*
	Em Go, a definição de público e privado é feita através de convenção.
	Inicio com letra Maiúscula- Publico
	Inicio com letra Minúscula- Privado

	A visibilidade é a nível de pacote, ou seja, uma váriavel privada, não é visível
	fora do pacote, mas pode ser acessada em outro arquivo dentro do pacote

	Por isso, não podemos ter funções duplicadas nos arquivos, pois a função
	é unica por pacote
*/

// Ponto representa uma coordenada do plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distância linear entre dois` pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))

}
