package main

import (
	"fmt"
	"strconv"
)

func toBinary(i int) string {
	i64 := int64(i)

	return strconv.FormatInt(i64, 2)
}

func main() {
	num := 10
	fmt.Println("Binary : ", toBinary(num))
}
