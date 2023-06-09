package main

import (
	"fmt"
	"net/http"
	"time"
)

func CorredorSync(a, b string) (vencedor string) {
    duracaoA := medirTempoDeResposta(a)
    duracaoB := medirTempoDeResposta(b)

    if duracaoA < duracaoB {
        return a
    }

    return b
}

func medirTempoDeResposta(URL string) time.Duration {
    inicio := time.Now()
    http.Get(URL)
    return time.Since(inicio)
}

var limiteDeDezSegundos = 10 * time.Second

func CorredorAsync(a, b string) (vencedor string, erro error) {
    return Configuravel(a, b, limiteDeDezSegundos)
}

func Configuravel(a, b string, tempoLimite time.Duration) (vencedor string, erro error) {
	// O que o select te permite fazer é aguardar múltiplos canais.
	// O primeiro a enviar um valor "vence" e o código abaixo do case é executado.
    select {
    case <-ping(a):
        return a, nil
    case <-ping(b):
        return b, nil
    case <-time.After(tempoLimite):
        return "", fmt.Errorf("tempo limite de espera excedido para %s e %s", a, b)
    }
}

func ping(URL string) chan bool {
    ch := make(chan bool)
    go func() {
        http.Get(URL)
        ch <- true
    }()
    return ch
}

func main() {
	fmt.Println("ola mundo")
}