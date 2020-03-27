package main

import "fmt"

func main() {
	var arrayInt [3]int
	arrayString := [3]string{"a", "b", "c"}

	var matriz [3][2]int

	fmt.Println(arrayInt)
	fmt.Println(arrayString)

	arrayInt[0] = 1
	arrayInt[1] = 2
	arrayInt[2] = 3

	for i := 0; i < len(arrayInt); i++ {
		fmt.Println(arrayInt[i])
	}

	fmt.Println(matriz)
}
