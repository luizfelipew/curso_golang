package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { //if com init tb suportado no switch
		fmt.Println("Ganhou!!!!!", i)
	} else {
		fmt.Println("Perdeu prayboy!!!", i)
	}
}
