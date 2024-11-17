package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	masukinName := Man{"Febry"}
	masukinName.Married()

	fmt.Println(masukinName.Name)
}
