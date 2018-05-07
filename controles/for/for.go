package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // nao existe while em GO
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando... tudo")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Ímpar ")
		}
	}

	// for infinito
	for {
		fmt.Println("Executando para sempre... ")
		time.Sleep(time.Second)
	}
}
