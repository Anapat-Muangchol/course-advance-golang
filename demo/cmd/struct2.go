package main

import "fmt"

type A1 struct {
	Id int
}

// A2 Composition over inheritance
type A2 struct {
	Name string
	A1
}

type ServiceA struct {
	A2
}

// NewServiceA pattern of init service, create function NewService for use
func NewServiceA(id int, name string) ServiceA {
	a1 := A1{Id: id}
	a2 := A2{A1: a1, Name: name}
	return ServiceA{A2: a2}
}

func RunStruct2() {
	fmt.Println("===== RunStruct2() =====")

	a1 := A1{}
	a1.Id = 2
	//a1.Name = "John"	// error not have this properties
	fmt.Println("a1 = ", a1)

	a2 := A2{}
	a2.Id = 2
	a2.Name = "Johnny"
	fmt.Println("a2 = ", a2)

	service := NewServiceA(3, "Anapat")
	fmt.Println("service = ", service)

	service.Name = "Anapat.Mua"
	fmt.Println("service = ", service)
}
