package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

/*
	Em Go, os arquivos de teste devem seguir a convenção de terminar com "_test"
	O nome antes do "_test" não precisa ser o mesmo do arquivo testado

	Assim como existe essa convenção para o nome do arquivo, existe também uma
	convenção para o nome do método de teste.
	Ele deve começar com a palavra "Test", não precisa ser o nome do método a
	ser testado depois de "Test".

	Outra convenção, é que precisa receber como paramêtro, um ponteiro para
	testing.T (t *testing.T), mesmo que ele não seja usado

	Podemos executar os testes via cli com o comando "go test".
	Caso esteja dentro da pasta, somente "go test", se estiver fora da pasta
	podemos executar com "go test ./..." que ele vai olhar para todos os testes
	dentro das subpastas
*/

func TestMedia(t *testing.T) {
	t.Parallel()
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
