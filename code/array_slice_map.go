package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := []string{"Hello", "world"}
	c := map[string]int{"valor1": 1, "valor2": 2}
	d := make([]int, 4)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println()

	d[2] = 5

	fmt.Println(a[2])
	fmt.Println(b[0], b[1])
	fmt.Println(c["valor1"])
	fmt.Println(d)
}
