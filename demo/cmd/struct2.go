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
func NewServiceAExam1(id int, name string) ServiceA {
	a1 := A1{Id: id}
	a2 := A2{A1: a1, Name: name}
	return ServiceA{A2: a2}
}

func NewServiceAExam2(id int, name string) ServiceA {
	serviceA := ServiceA{}
	serviceA.Id = id
	serviceA.Name = name
	return serviceA
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

	service1 := NewServiceAExam1(3, "Anapat")
	fmt.Println("service1 = ", service1)

	service1.Name = "Anapat.Mua"
	fmt.Println("service1 = ", service1)

	service2 := NewServiceAExam2(4, "Anapat")
	fmt.Println("service2 = ", service2)

	service2.Name = "Anapat.Muangchol"
	fmt.Println("service2 = ", service2)
}
