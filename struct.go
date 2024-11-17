package main

import "fmt"

type Customers struct {
	Name    string
	Address string
	Age     int
}

func (customers Customers) sayHello(name string) {
	fmt.Println("Hello ", name, "my name is", customers.Name)
}

func main() {
	var data Customers
	data.Name = "Belisca"
	data.Address = "Lampung Tengah"
	data.Age = 21

	fmt.Println(data)
	fmt.Println(data.Name)
	fmt.Println(data.Age)
	fmt.Println(data.Address)

	data.sayHello("Agus Bangunan")
}
