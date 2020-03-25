package main

import "fmt"

func main() {
	/*
	   Es una dirección de memoria
	   Valor zero nil
	*/

	var x, y *int
	entero := 5

	x = &entero
	y = &entero

	*x = 6

	fmt.Println(*x)
	fmt.Println(*y)

}
