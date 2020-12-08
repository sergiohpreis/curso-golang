package main

import "fmt"

func inc1(n int) {
	n++
}

/*
Um ponteiro em go é representado por um *, no caso
dessa função, ela recebe um ponteiro de inteiro como argumento
*/
func inc2(n *int) {
	/*
		O * é usado para desreferenciar, ou seja, para ter
		acesso ao valor no qual o ponteiro aponta
	*/
	*n++ // acessa o valor do ponteiro e incrementa
}

func main() {
	n := 1

	/*
		Nesse caso, o valor da váriavel "n" não é modificado
		pois "inc1" usa o valor da váriavel, e não sua
		refêrencia
	*/
	fmt.Println(n) // 1

	/*
		Nesse caso, o valor da váriavel "n" é modificado
		pois "inc2" usa a referência da váriavel, e não seu valor
	*/
	inc2(&n)       // & usado para obter endereço da variável
	fmt.Println(n) // 20

	/*
		O cenário ideal é sempre manter a variável com o menor escopo
		possível.
		Passar a referência de uma váriavel para outros metodos pode
		causar efeitos colaterais indesejados na aplicação, pois a função
		deixa de ser pura, aumentando a probabilidade de bugs por conta dos efeitos
		colaterais

		PRIORIZE O USO DO VALOR DE UMA VARIAGEM AO INVÉS DE SUA REFERÊNCIA
		PARA NÃO PERDER O CONTROLE DO CÓDIGO
	*/
}
