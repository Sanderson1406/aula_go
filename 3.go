package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var base int8
	var altura int8
	var profun int8
	fmt.Print("Qual a base da caixa?")
	fmt.Scan(&base)
	fmt.Print("Qual a altura da caixa?")
	fmt.Scan(&altura)
	fmt.Print("Qual a profundidade da caixa?")
	fmt.Scan(&profun)
	fmt.Println("O colume da caixa é", base*altura*profun)

}
