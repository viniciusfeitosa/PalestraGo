package main

import "fmt"

func main() {
	x := 1

	switch x {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("Um")
	case 2:
		fmt.Println("Dois")
	default:
		fmt.Println("Não sei o número")

	}
}
