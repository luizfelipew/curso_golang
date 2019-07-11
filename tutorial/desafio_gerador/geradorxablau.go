package main

import (
	"fmt"
	"math/rand"
	"time"
)

func geradorxablau() []string {
	primes := []string{"quêêêêêê",
		"perc",
		"fiforn",
		"tipoType",
		"énada",
		"percAss",
		"vouBotarNaSuaBranch",
		"tranquera",
		"canalha",
		"xablau",
		"gcp",
		"putz",
		"chorando",
		"valValor",
		"trocaBem",
		"caaaraaamba",
		"taCerto",
		"QuatroEvinte",
		"naoSePodeGanharTodas",
		"aceiteMinhaDerrotaPoisEuJaAceitei"}

	return shuffle(primes)
}

func shuffle(vals []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(vals))
	n := len(vals)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals))
		ret[i] = vals[randIndex]
		vals = append(vals[:randIndex], vals[randIndex+1:]...)
	}
	return ret
}

func main() {
	fmt.Println(geradorxablau())
}
