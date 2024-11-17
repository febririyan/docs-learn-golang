package main

import "fmt"

func sayHelloTo(fristName string, lastName string) {
	fmt.Println("Hello ", fristName, lastName)
}

func main() {
	sayHelloTo("Febri", "Riyanto")
	sayHelloTo("Berlian", "Puspita")
}
