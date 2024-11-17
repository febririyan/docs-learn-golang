package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := &address1 // kita membuat pointer

	address2.Province = "Lampung"
	fmt.Println(address1) // Datanya akan berubah karena address 2 itu pointer dari address 1
	fmt.Println(address2) // Datanya berubah lampung

	//address2 = &Address{"DKI", "Bandung", "Medan"}

	*address2 = Address{"DKI", "Bandung", "Medan"}
	fmt.Println(address1)
	fmt.Println(address2)
}