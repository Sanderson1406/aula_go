package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var dolaremr float32
	var dolar float32
	fmt.Print("Quanto você tem em dolar?")
	fmt.Scan(&dolar)
	fmt.Print("Quanto é dolar em reais?")
	fmt.Scan(&dolaremr)
	fmt.Println("Você tem em reais", dolar*dolaremr)
}
