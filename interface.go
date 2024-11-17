package main

import "fmt"

type HasName interface {
	getName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

func main() {
	person := Person{Name: "Febry"}
	SayHello(person)
}
