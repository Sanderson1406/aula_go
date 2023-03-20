package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var base int8
	var altura int8
	var result int8
	var area int8
	fmt.Print("Qual a base do triângulo?")
	fmt.Scan(&base)
	fmt.Print("Qual a altura do triângulo?")
	fmt.Scan(&altura)
	fmt.Println(base * altura)
	fmt.Scan(&result)
	fmt.Scan(&area)
	fmt.Print(area)

}
