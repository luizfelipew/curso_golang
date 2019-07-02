package main

import (
	"fmt"
	"strconv"
)

func toBinary(i int) string {
	i64 := int64(i)

	return strconv.FormatInt(i64, 2)
}

func toBinary2(i int) string {

	binario := ""

	num := strconv.Itoa(i)
	//os.Args[i]

	for len(num)-1 > 0 {
		fmt.Println("passou aqui")
		if i%2 == 0 {
			binario = "0" + binario
		} else {
			binario = "1" + binario
		}
	}
	return binario
}

func main() {
	num := 2
	// fmt.Println("Binary : ", toBinary(num))
	fmt.Println("Binary : ", toBinary2(num))
}
