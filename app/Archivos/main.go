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

		 r := recover()
		 if r != nil{
		 	fmt.Println("recover from ",r)
		 }
	}()

	if err != nil{
		panic(err)
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan(){
		linea := scanner.Text()
		fmt.Println(linea)
	}

	return true
}
