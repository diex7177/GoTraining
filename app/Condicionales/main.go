package main

import "fmt"

/*
 == igual a
 != diferente de
 > mayor que
 < menor que
 >= mayor o igual
 <= menor o igual
 && AND
 || OR
*/

func main() {
	y := 9
	x := 9

	if x > y {
		fmt.Printf("%d es mayor que %d \n", x, y)
	} else if y > x {
		fmt.Printf("%d es mayor que %d \n", y, x)
	} else {
		fmt.Printf("%d es iqual que %d", x, y)
	}
}
