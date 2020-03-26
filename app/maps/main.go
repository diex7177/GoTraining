package main

import "fmt"

func main() {
	x := make(map[int]string, 7)

	x[1] = "Lunes"
	x[2] = "Martes"
	x[3] = "Miercoles"
	x[4] = "Jueves"
	x[5] = "Viernes"
	x[6] = "Sábado"
	x[7] = "Domingo"

	fmt.Println(x)
	fmt.Println(x[1])

	edades := map[string]int{
		"Ana":   18,
		"Diego": 29,
		"Peter": 35,
	}

	fmt.Println(edades)

	delete(edades, "Diego")

	fmt.Println(edades)

	if edad, exits := edades["Ana"]; exits {
		if edad >= 18 {
			fmt.Println("Es Mayor de edad")
		} else {
			fmt.Println("Es menor de edad")
		}

	} else {
		fmt.Println("El dato no existe")
	}

	fmt.Println("Tamaño del map ", len(edades))

	for nombre, edad := range edades {
		fmt.Printf("la edad de %s es: %d \n", nombre, edad)
	}
}
