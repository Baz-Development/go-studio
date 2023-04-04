package main

import (
	"fmt"
)

// Estrutura da função parametro
type VerificadorWebsite func(string) bool
// Estrutura do canal
type resultado struct {
    string
    bool
}

func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	// Instancia um map de resultado
    resultados := make(map[string]bool)
	// Instancia um canal para que não aconteça deadlocks na inserção de dados no map
	canalResultado := make(chan resultado)

    for _, url := range urls {
		// Inicia uma nova go routine para cada loop
		// Ao passar u como parametro, fixa o valor recebino não alterando ele durante o tempo de runtime
        go func(u string) {
			// Insere a informação pelo canal 
			canalResultado <- resultado{u, vw(u)}
        }(url)
    }

    for i := 0; i < len(urls); i++ {
		// Atribui as informações do canal em uma variavel auxiliar
        resultado := <-canalResultado
		// Atribui as informações da variavel auxiliar para o map de retorno
        resultados[resultado.string] = resultado.bool
    }

    return resultados
}

func main() {
	fmt.Println("ola mundo")
}