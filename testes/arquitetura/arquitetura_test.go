package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() // roda os testes em paralelo, pode trazer ganho de performance
	if runtime.GOARCH == "amd64" {
		// pula o teste e não executa
		t.Skip("Não funciona em arquitetura amd64")
	}
	// ...
	t.Fail()
}
