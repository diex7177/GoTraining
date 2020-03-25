package main

import "fmt"

func main() {
	slice := make([]int, 0)
	fmt.Println(slice)
	slice = append(slice, 2)
	fmt.Println(slice)
	fmt.Println("Lenght", len(slice))
	fmt.Println("capacity", cap(slice))
}
