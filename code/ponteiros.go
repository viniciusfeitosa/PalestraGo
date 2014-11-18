package main

import "fmt"

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

	y := new(int)
	um(y)
	fmt.Println(*y)
}

func zero(x *int) {
	*x = 0
}

func um(y *int) {
	*y = 1
}
