package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado .....
	fmt.Println("Teste" + string(97))

	// inteiro para string
	fmt.Println("Teste2 " + strconv.Itoa(123))

	// string para int = ele retorna dois valores o primeiro e um numero o segundo eh um erro
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("1") // or true or number 1
	if b {
		fmt.Println("Verdadeiro")
	}

}
