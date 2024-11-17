package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("for di ulang lagi", counter)
		counter++
	}

	fmt.Println("selesai")

	name := []string{"Berlian", "Febry", "Riyan", "Berliana"} // akses manual
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	for index, name := range name {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range name {
		fmt.Println(name)
	}
}
