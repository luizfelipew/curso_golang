package main

import (
	"fmt"
)

func main() {
	i := 1

	// GO não tem aritmática de ponteiros
	// p++
	var p *int = nil

	// endereco de memoria da variavel i para o p '&'
	p = &i
	*p++ // desreferenciando o poteiro par apegar o valor
	i++

	fmt.Println(p, *p, i, &i)
}
