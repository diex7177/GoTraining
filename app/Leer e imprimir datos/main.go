package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	edad := 22
	nombre := "Diego"
	flag := true
	salario := 14.3
	var nombreUsuario string

	fmt.Printf("Mi edad es %d \n", edad)
	fmt.Printf("Mi nombre es %v \n", nombre)
	fmt.Printf("Soy el mejor %t \n", flag)
	fmt.Printf("Mi salario es %f \n", salario)

	fmt.Println("Cual es su nombre: ")
	fmt.Scanf("%v\n", &nombreUsuario)
	fmt.Printf("Su nombre es %v \n", nombreUsuario)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu apellido: ")
	apellido, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Ocurrio un error: ", err)
	} else {
		fmt.Printf("Su apellido es: %v \n", apellido)
	}
}
