package main

import (
	"fmt"
	"strconv"
)

func toBinary(i int) string {
	i64 := int64(i)

	return strconv.FormatInt(i64, 2)
}

func toBinary2(number int) string {

	binario := ""
	num := int64(number)

	//num := strconv.Itoa(number)
	//os.Args[i]
	// fmt.Printf("Teste: %d ", len(num))
	for i := 0; num > 0; i++ {

		if num == 0 {
			return "0"
		}
		if num%2 == 0 {
			binario = "0" + binario
		} else {
			binario = "1" + binario
		}
		num /= 2
	}

	return binario
}

func main() {
	num := 10
	// fmt.Println("Binary : ", toBinary(num))
	fmt.Println("Binary : ", toBinary2(num))
}
