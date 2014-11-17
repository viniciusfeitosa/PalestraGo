package main

import "fmt"

func main() {
	soma := func(x, y int) int {
		return x + y
	}
	fmt.Println(soma(1, 5))
}
