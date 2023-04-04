package main

import (
	"fmt"
	"sync"
)

type Contador struct {
    mu sync.Mutex
    valor int
}

func (c *Contador) Incrementa() {
	// Qualquer goroutine que chamar a func Incrementa será bloqueada até o fim do processo de incrementação
    c.mu.Lock()
	// defer é um prefixo que faz com que o codigo sejá execultado somente no fim de todos os processos da função
    defer c.mu.Unlock()
    c.valor++
}

func (c *Contador) Valor() int {
	return c.valor
}

func main() {
	fmt.Println("ola mundo")
}