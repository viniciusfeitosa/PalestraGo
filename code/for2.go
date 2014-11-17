package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i += 1
	}

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for y, z := range x {
		fmt.Printf("Index: %d Value: %d \n", y, z)
	}

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range a {
		fmt.Printf("Value: %d \n", v)
	}
}
