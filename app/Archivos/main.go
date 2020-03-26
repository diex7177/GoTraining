package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	isCorrect := readFile()
	fmt.Println(isCorrect)
}

func readFile() bool{
	archivo, err := os.Open("./test.txt")
	defer func(){
		 archivo.Close()
		 fmt.Println("Difer")
	}()

	if err != nil{
		fmt.Println("Hubo un error ", err)
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan(){
		linea := scanner.Text()
		fmt.Println(linea)
	}

	return true
}
