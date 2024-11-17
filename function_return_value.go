package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	result := getHello("Febry")
	fmt.Println(result)

	fmt.Println(getHello("Berliana"))
	fmt.Println(getHello("Riyan"))
}
