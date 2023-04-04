package main

import (
	"reflect"
	"testing"
	"time"
)

func mockVerificadorWebsite(url string) bool {
	if url == "https://baz.dev.br" {
		return false
	}
	return true
}

func TestVerificadorWebsite(t *testing.T) {
	websites := []string{
        "http://google.com",
        "https://github.com/felipe-baz",
        "https://baz.dev.br",
    }

	esperado := map[string]bool {
		"http://google.com": true,
		"https://github.com/felipe-baz": true,
		"https://baz.dev.br": false,
	}

	resultado := VerificaWebsites(mockVerificadorWebsite, websites)

	if !reflect.DeepEqual(esperado, resultado) {
		t.Errorf("resultado %v, esperado %v", resultado, esperado)
	}
}

func slowStubVerificadorWebsite(_ string) bool {
    time.Sleep(20 * time.Millisecond)
    return true
}

// Função para testes em massa
func BenchmarkVerificaWebsites(b *testing.B) {
    urls := make([]string, 1000)
    for i := 0; i < len(urls); i++ {
        urls[i] = "uma url"
    }

    for i := 0; i < b.N; i++ {
        VerificaWebsites(slowStubVerificadorWebsite, urls)
    }
}