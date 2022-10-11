package main

import (
	"fmt"
	"unsafe"
)

type MyStruct struct {
	b bool    // 1 Byte
	f float64 // 8 Byte
	i int32   // 4 Byte
}

type MyStructReduceResource struct {
	f float64 // 8 Byte
	i int32   // 4 Byte
	b bool    // 1 Byte
}

func RunStructReduceResource() {
	fmt.Println("===== RunStructReduceResource() =====")

	s1 := MyStruct{}
	fmt.Printf("s1 use resource %d bytes\n", unsafe.Sizeof(s1)) // 24 bytes

	s2 := MyStructReduceResource{}
	fmt.Printf("s2 use resource %d bytes\n", unsafe.Sizeof(s2)) // 16 bytes
}
