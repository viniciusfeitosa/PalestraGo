package main

import "fmt"

func main() {
	a := []float64{1.2, 4.0, 3.5}
	fmt.Println(media(a))
	fmt.Println(retornoDuplo())
}

func media(xs []float64) float64 {
	var x float64
	for _, y := range xs {
		x += y
	}
	return x / float64(len(xs))
}

func retornoDuplo() (a string, b string) {
	a, b = "Olá", "mundo"
	return a, b
}
