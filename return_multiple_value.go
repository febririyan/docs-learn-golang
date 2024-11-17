package main

import "fmt"

func getFullName() (string, string) {
	return "Febry", "Riyanto"
}

func main() {
	//fristName, lastName := getFullName()
	//fmt.Println(fristName, lastName)

	// Jika menghiraukan salah satu

	fristName, _ := getFullName()
	fmt.Println(fristName)
}
