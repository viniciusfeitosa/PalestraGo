package main

import "fmt"

var (
	a string
	b int64 = 2
)

var c float64 = 1.0

const (
	A string  = "A"
	B float64 = 1.1
)

const C string = "D"

func main() {
	a = "1"
	d := "teste"
	fmt.Println(a, b, c, d)
	fmt.Println(A, B, C)
}
