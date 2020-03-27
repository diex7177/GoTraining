package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go nombreLento("Diego")
	fmt.Println("Que aburrido")
	var wait string
	fmt.Scanln(&wait)
}

func nombreLento(nombre string) {
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
