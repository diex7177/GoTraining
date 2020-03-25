package main

import "fmt"

type Human struct {
	name string
}

type Tutor struct {
	Human
}

func (tutor *Tutor) hablar() string {
	return tutor.Human.hablar() + " Bienvenido"
}

func (human *Human) hablar() string {
	return "bla bla bla..."
}

func main() {
	tutor := Tutor{Human{"Diego"}}

	fmt.Println(tutor.name)
	fmt.Println(tutor.hablar())
}
