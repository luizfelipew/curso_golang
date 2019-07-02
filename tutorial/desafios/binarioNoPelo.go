package main

func toBinary2(i int) string {

	binario := ""

	num := int64(i)

	for num > 0 {
		if i%2 == 0 {
			binario = "0" + binario
		} else {
			binario = "1" + binario
		}
	}

	return binario
}

// func main() {
// 	num := 10
// 	fmt.Println("Binary : ", toBinary2(num))
// }
