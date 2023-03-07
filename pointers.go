package main

import "fmt"

func main() {
	a := 10
	b := &a         // Endereço de memória
	fmt.Println(*b) // VALOR do endereço de memória
	fmt.Println(b)  // Endereço de memória
}
