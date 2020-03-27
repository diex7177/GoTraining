package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func main() {
	user := User{edad: 29, nombre: "Diego", apellido: "Ospina"}

	user1 := new(User)

	user1.nombre = "Lakshmi"

	fmt.Println(user)
	fmt.Println(user1)
}
