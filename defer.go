package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runAplication() {

	defer logging()

	fmt.Println("Run aplication")
}

func main() {
	runAplication()
}
