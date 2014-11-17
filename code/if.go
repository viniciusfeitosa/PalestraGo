package main

import "fmt"

func main() {

	x := 1

	if x == 1 {
		fmt.Println("Igual a um")
	} else if x == 2 {
		fmt.Println("Igual a dois")
	} else {
		fmt.Println("Não sei o valor")
	}

	if y := 2; y == 2 {
		fmt.Println("Igual a dois")
	}
}
