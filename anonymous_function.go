package main

import "fmt"

type Backlist func(string) bool

func registerUser(name string, backlist Backlist) {
	if backlist(name) {
		fmt.Println("Kamu di backlist", name)
	} else {
		fmt.Println("Selamat Datang", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Febry", blacklist)
}
