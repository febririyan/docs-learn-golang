package main

import "fmt"

func main() {
	name := "Berlian"

	switch name {
	case "Adi":
		fmt.Println("Hello Adi")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hello World")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Terlalu Benar")
	}
}
