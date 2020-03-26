package main

import "fmt"

func print(cadena string) {
	fmt.Println(cadena)
}

func main() {
	cadena := "Hola Mundo"

	imprimir := print

	imprimir(cadena)

	func2 := func() {
		fmt.Println(cadena)
	}

	func2()
}
