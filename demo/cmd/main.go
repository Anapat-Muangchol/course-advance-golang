package main

import (
	"demo"
	"fmt"
)

func main() {
	m, err := demo.SayHi()
	if err != nil {
		fmt.Println(m)
	}
}
