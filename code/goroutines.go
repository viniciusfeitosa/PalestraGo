package main

import (
	"fmt"
	"runtime"
)

func f(n string) {
	for i := 0; i < 100; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go f("ping")
	go f("pong")
	var input string
	fmt.Scanln(&input)
}
