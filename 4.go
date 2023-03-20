package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var n1 int8
	var n2 int8
	var n3 int8
	var n4 int8
	fmt.Print("Qual é o primeiro número")
	fmt.Scan(&n1)
	fmt.Print("Qual é o segundo número")
	fmt.Scan(&n2)
	fmt.Print("Qual é o terceiro número")
	fmt.Scan(&n3)
	fmt.Print("Qual é o quarto número")
	fmt.Scan(&n4)
	fmt.Println("A sua média aritmética é", n1*n2*n3*n4/2)

}
