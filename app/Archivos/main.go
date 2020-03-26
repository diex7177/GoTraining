package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	archivo, err := os.Open("./test.txt")

	if err != nil{
		fmt.Println("Hubo un error ", err)
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan(){
		linea := scanner.Text()
		fmt.Println(linea)
	}
}
