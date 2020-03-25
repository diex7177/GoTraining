package main

import (
	"fmt"
	"strconv"
)

func main()  {
	edad  := 23
	mesNacimiento := "15"
	mesNacimientoInt,_ := strconv.Atoi(mesNacimiento)

	fmt.Println("Edad: " + strconv.Itoa(edad))
	fmt.Println(mesNacimientoInt)
}