package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	t.Run("compara a velocidade de servidores, retornando o endereço do mais rápido", func(t *testing.T) {
		// Gera um servidor Fake
        servidorLento := criarServidorComAtraso(20 * time.Millisecond)
        servidorRapido := criarServidorComAtraso(0 * time.Millisecond)

		// Fecha as conexões
		// Ao chamar uma função com o prefixo defer, ela será chamada após o término da função que a contém.
        defer servidorLento.Close()
        defer servidorRapido.Close()

		// Coleta as url dos servidores fakes
        URLLenta := servidorLento.URL
        URLRapida := servidorRapido.URL

        esperado := URLRapida
        resultado, err := CorredorAsync(URLLenta, URLRapida)

        if err != nil {
            t.Fatalf("não esperava um erro, mas obteve um %v", err)
        }

        if resultado != esperado {
            t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
        }
    })

    t.Run("retorna um erro se o servidor não responder dentro de 10s", func(t *testing.T) {
		// Gera um servidor Fake
        servidor := criarServidorComAtraso(25 * time.Millisecond)

		// Fecha as conexões
		// Ao chamar uma função com o prefixo defer, ela será chamada após o término da função que a contém.
        defer servidor.Close()

        _, err := Configuravel(servidor.URL, servidor.URL, 20*time.Millisecond)

        if err == nil {
            t.Error("esperava um erro, mas não obtive um")
        }
    })
}

func criarServidorComAtraso(atraso time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(atraso)
        w.WriteHeader(http.StatusOK)
    }))
}