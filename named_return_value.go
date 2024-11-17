package main

import "fmt"

func getComplateName() (fristName, middleName, lastName string) {
	fristName = "Febry"
	middleName = "Riyan"
	lastName = "Hidayat"

	return fristName, middleName, lastName
}

func main() {
	a, b, c := getComplateName()
	fmt.Println(a, b, c)
}
