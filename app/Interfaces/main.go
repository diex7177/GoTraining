package main

import "fmt"

type User interface {
	Permissions() int //1 - 5
	Nombre() string
}

type Admin struct {
	name string
}

func (admin Admin) Nombre() string {
	return admin.name
}

func (admin Admin) Permissions() int {
	return 5
}

type Editor struct {
	name string
}

func (editor Editor) Nombre() string {
	return editor.name
}

func (editor Editor) Permissions() int {
	return 3
}

func Auth(user User) string {
	if user.Permissions() == 5 {
		return user.Nombre() + " Tiene permisos de administrador"
	} else if user.Permissions() == 3 {
		return user.Nombre() + " Tiene permisos de editor"
	} else {
		return "no tiene permisos de administrador"
	}
}

func main() {
	usuarios := []User{Admin{"Diego"}, Editor{"Andres"}}

	for _, usuario := range usuarios {
		fmt.Println(Auth(usuario))
	}
}
