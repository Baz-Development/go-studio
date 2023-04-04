package main

import (
	"fmt"
	"reflect"
)

func percorre(x interface{}, fn func(entrada string)) {
	valor := obtemValor(x)

	percorreValor := func(valor reflect.Value) {
        percorre(valor.Interface(), fn)
    }

	// Verifica o tipo do campo
	switch valor.Kind() {
    case reflect.String:
        fn(valor.String())
    case reflect.Struct:
		// O valor tem um metodo chamado NumField que retorna a quantidade de campos do valor
        for i := 0; i < valor.NumField(); i++ {
			// extrai-se o valor das sequintes formas: valor.Field[campo] ou valor.Index[índice]
            percorreValor(valor.Field(i))
        }
    case reflect.Slice, reflect.Array:
        for i := 0; i < valor.Len(); i++ {
			// extrai-se o valor das sequintes formas: valor.Field[campo] ou valor.Index[índice]
            percorreValor(valor.Index(i))
        }
    case reflect.Map:
        for _, chave := range valor.MapKeys() {
            percorreValor(valor.MapIndex(chave))
        }
    }
}

func obtemValor(x interface{}) reflect.Value {
    // reflection verifica as propriedades de x
	valor := reflect.ValueOf(x) // ValueOf retorna um Value (valor) de determinada variável. 

	// Não é possível usar o NumField em um ponteiro e precisamos extrair o valor antes disso usando Elem().
	if valor.Kind() == reflect.Ptr {
        valor = valor.Elem()
    }

    return valor
}

func main() {
	fmt.Println("ola mundo")
}