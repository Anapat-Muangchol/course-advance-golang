package main

import (
	"demo"
	"fmt"
)

func sayHi() {
	m, err := demo.SayHi()
	if err == nil {
		fmt.Println(m)
	}
}

func loop() {
	// Loop :: for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Loop :: while
	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Infinity Loop
	//for {}

	m, err := demo.SayHi()
	if err == nil {
		fmt.Println(m)
	}
}

func testDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("End")
}
