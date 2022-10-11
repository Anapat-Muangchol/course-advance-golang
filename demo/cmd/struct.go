package main

import "fmt"

// Class
//  + Properties
//  + Behavior

type DataInterface interface {
	process() string
}

type Data struct {
	Id   int
	Name string
}

func (d Data) process() string {
	return fmt.Sprintf("Call process with id=%d", d.Id)
}

// like a toString() method in Java
func (d Data) String() string {
	return fmt.Sprintf("[Id: %d, Name: %s]", d.Id, d.Name)
}

func printString(b DataInterface) {
	fmt.Println(b.process())
}

func RunStruct() {
	fmt.Println("===== RunStruct() =====")

	d1 := Data{1, "John"}
	d2 := Data{Id: 2, Name: "Johnny"}
	fmt.Println(d1, d2)

	fmt.Println(d2.process())
	fmt.Println(d2.Name)

	printString(&d2)
}
