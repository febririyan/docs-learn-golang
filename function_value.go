package main

import "fmt"

func getGodBye(name string) string {
	return "Hello " + name
}

func main() {
	helloTeman := getGodBye

	fmt.Println(helloTeman("Febry"))
}
