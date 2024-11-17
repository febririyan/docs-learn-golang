package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Febry",
		"age":     "25",
		"address": "Lampung Tengah",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Febry"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
