package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	copia := make([]int, 5)
	slice = append(slice, 1)

	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)
}
