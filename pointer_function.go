package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Kores Selatan"
}

func main() {
	var address = Address{}
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}
