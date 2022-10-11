package main

import (
	"fmt"
	"io/ioutil"
	"runtime/debug"
)

func RunPanic() {
	defer panicHandler()

	// Read data from file
	b, err := ioutil.ReadFile("try_panic.txt")
	if err != nil {
		// log.Fatalf("error %v", err)	// ไม่ควรใช้ log.Fatal เนื่องจาก จะ exit program ทันที
		panic("read file error")
	}

	fmt.Println(string(b))
}
func panicHandler() {
	err := recover()
	if err == "read file error" {
		fmt.Println("Try to recover from panic")
		// TODO
		debug.PrintStack()
	}
}
