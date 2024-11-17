package main

import (
	"fmt"
	"learn-go/database"
	"learn-go/helper"
)

func main() {
	result := helper.NameDepan("Febri")
	fmt.Println(result)

	fmt.Println(helper.Aplication)

	fmt.Println(database.GetDatabase())
}
