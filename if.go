package main

import "fmt"

func main() {
	name := "Budibudibudi"

	if name == "Febry" {
		fmt.Println("Yes i am febry")
	} else if name == "Berlian" {
		fmt.Println("Yes i am berlian")
	} else {
		fmt.Println("Nama tidak cocok")
	}

	if length := len(name); length > 8 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama benar bagus!")
	}
}
