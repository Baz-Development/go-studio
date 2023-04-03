package main

import (
	"errors"
	"fmt"
)

// Define um type chamado bitcoin que é inteiro
type Bitcoin int

// Define uma interface para impressão do bitcoin
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}

// Define objeto Carteira
type Carteira struct {
	saldo Bitcoin
}

// Define o metodo Depositar
func (c *Carteira) Depositar(quant Bitcoin) {
	fmt.Printf("O endereço do saldo no Depositar é %v \n", &c.saldo)
	c.saldo += quant
}

// Define o metodo Sacar
func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

// Cria uma constante com a mensagem para saldo insuficiente
var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

// Define o metodo Sacar
func (c *Carteira) Sacar(quant Bitcoin) error {
	fmt.Printf("O endereço do saldo no Sacar é %v \n", &c.saldo)

	if quant > c.saldo {
        return ErroSaldoInsuficiente
    }

    c.saldo -= quant
    return nil
}

func main() {
}