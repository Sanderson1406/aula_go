package main

import "fmt"

func main() {
	var name string
	var id uint
	var peso float64
	fmt.Print("Qual é o seu nome? ")
	fmt.Scan(&name)
	fmt.Print("Qual a sua idede? ")
	fmt.Scan(&id)
	fmt.Print("Qual seu peso? ")
	fmt.Scan(&peso)
	fmt.Println("Olá", name, ", você tem", id, "anos e pesa", peso, "...")
}
